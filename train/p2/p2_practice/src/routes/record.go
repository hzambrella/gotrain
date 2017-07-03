package routes

import (
	"encoding/json"
	"engine/datastore"
	"errors"
	"fmt"
	"model"
	"net/http"
	"strconv"

	"lib/bs"
	"lib/cloudprt"
	"lib/loghz"
)

var d datastore.Data

func PrtOrder(w http.ResponseWriter, r *http.Request) {
	// 错误信息
	var errMsg string
	// 向前端发送的信息
	var result []byte
	// 打印机类别
	var prtType int = cloudprt.S1

	d, err := datastore.ParseDataFromFile("key.cfg")
	if err != nil {
		loghz.Error(err)
	}

	defer func() {
		if len(errMsg) != 0 {
			result = bs.S2B(errMsg)
		}
		if _, err := w.Write(result); err != nil {
			loghz.Error(err)
		}
	}()

	recIdStr := r.FormValue("recId")
	if recIdStr == "" {
		errMsg = "recId is nil"
		loghz.Error(errors.New(errMsg))
		return
	}

	if len(d["key"]) == 0 {
		errMsg = "key is nil"
		loghz.Error(errors.New(errMsg))
		return
	}

	prt, err := cloudprt.NewPrt(prtType, d["deviceNo"], d["key"])
	if err != nil {
		errMsg = err.Error()
		loghz.Error(errors.New(errMsg))
		return
	}

	//db
	orderDB, err := model.NewOrderDB(d["dsn"])
	if err != nil {
		errMsg = err.Error()
		loghz.Error(errors.New(errMsg))
		return
	}

	recId, err := strconv.ParseUint(recIdStr, 10, 64)
	if err != nil {
		errMsg = err.Error()
		loghz.Error(errors.New(errMsg))
		return
	}

	//get record
	record, err := orderDB.GetRecordById(recId)
	if err != nil {
		errMsg = err.Error()
		loghz.Error(errors.New(errMsg))
		return
	}

	//get table
	table, err := orderDB.GetTableById(record.TableId)
	if err != nil {
		errMsg = err.Error()
		loghz.Error(errors.New(errMsg))
		return
	}

	shop, err := orderDB.GetShopById(record.ShopId)
	if err != nil {
		errMsg = err.Error()
		loghz.Error(errors.New(errMsg))
		return
	}

	// 定义这个的原因是当时前端也需要数据，没设计好，写的不好
	var foods []FoodList
	// 云打印机需要的数据
	var goodPrt []*cloudprt.Goods
	// 总价
	var totalPrice float64

	// data
	for k, v := range record.GoodNum {
		var food FoodList
		food.Num = v
		food.Id = record.GoodId[k]
		// TODO:Price type change to money
		food.UintPrice = record.GoodHisPrice[k]
		//注意！！food.Amount单个食物的总价,单价乘以数量！！
		food.Amount = record.GoodHisPrice[k] * float64(v)
		food.Name = record.GoodHisName[k]
		foods = append(foods, food)
		goodPrt = append(goodPrt, &cloudprt.Goods{food.Name, food.UintPrice, food.Num})
		totalPrice = totalPrice + food.Amount
	}

	var eatType string
	switch record.EatType {
	case 0:
		eatType = "堂食"
	case 1:
		eatType = "带走"
	}

	// 定义这个的原因是当时前端也需要数据，没设计好，写的不好
	menu := MenuList{
		ShopName:   shop.Name,
		TableName:  table.Name,
		RecId:      recId,
		EatType:    eatType,
		FoodsList:  foods,
		Memo:       record.Memo,
		TotalPrice: totalPrice,
		CreateTime: record.CreateTime,
	}

	bf, err := json.Marshal(&foods)
	if err != nil {
		errMsg = err.Error()
		loghz.Error(errors.New(errMsg))
		return
	}

	for _, v := range bf {
		result = append(result, v)
	}

	bm, err := json.Marshal(&menu)
	if err != nil {
		errMsg = err.Error()
		loghz.Error(errors.New(errMsg))
		return
	}

	for _, v := range bm {
		result = append(result, v)
	}

	shopMess := []string{"桌号", menu.TableName, "订单号", strconv.Itoa(int(recId)), "就餐方式", eatType}

	menuList := &cloudprt.MenuList{menu.ShopName, shopMess, []string{"备注", record.Memo}, goodPrt, record.CreateTime}
	xiaopiao := cloudprt.NewDefaultxiaoPiao(menuList)
	prtStr := xiaopiao.FormatToReciept()
	fmt.Println(prtStr)

	//不要浪费纸张，有把握了再打印
	if 1 == 0 {
		respMsg, err := prt.PrintString(prtStr)

		if err != nil {
			errMsg = err.Error()
			loghz.Error(errors.New(errMsg))
			return
		}

		var br []byte
		if respMsg != nil {
			loghz.Println(*respMsg)
			br, err = json.Marshal(&respMsg)
			if err != nil {
				errMsg = err.Error()
				loghz.Error(errors.New(errMsg))
				return
			}
		}

		for _, v := range br {
			result = append(result, v)
		}
	}

}

//TODO 查询打印机状态,json格式的结果返回到浏览器上
