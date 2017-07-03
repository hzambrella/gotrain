package main

import (
	"cloudprt"
	"strconv"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"model"
	"net/http"

	"engine/datastore"
	"lib/bs"
)

var (
	prtType  int = 0
	deviceNo string
	key      string
)

var d datastore.Data

func init() {
	var err error
	d, err = datastore.ParseDataFromFile("key.cfg")
	if err != nil {
		log.Fatal(err)
	}
	deviceNo = d["deviceNo"]
	key = d["key"]
	log.Println(d)
}

func main() {
	http.HandleFunc("/clouldprt/prtorder", prtOrder)
	addr := ":8080"
	fmt.Println("run at " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func prtOrder(w http.ResponseWriter, r *http.Request) {
	var errMsg string
	var result []byte
	var msg string

	defer func() {
		if len(errMsg) != 0 {
			result = bs.S2B(errMsg)
		}
		if _, err := w.Write(result); err != nil {
			log.Fatal(err)
		}
	}()

	shopId := r.FormValue("shopId")
	if len(key) == 0 {
		errMsg = "key is nil"
		log.Fatal(errors.New(errMsg))
		return
	}

	msg = r.FormValue("msg")
	if len(msg) == 0 {
		msg = "hello world"
	}

	prt, err := cloudprt.NewPrt(prtType, deviceNo, key)
	if err != nil {
		errMsg = err.Error()
		log.Fatal(errMsg)
		return
	}

	orderDB, err := model.NewOrderDB(d["dsn"])
	if err != nil {
		errMsg = err.Error()
		log.Fatal("haha"+errMsg)
		return
	}

	shopIdStr, err := strconv.ParseUint(shopId, 10, 0)
	if err != nil {
		errMsg = err.Error()
		log.Fatal(err.Error())
		return
	}

//TODO:order printOB
//改写下面代码，打印订单
	items, err := orderDB.GetItemsByShopId(shopIdStr)
	if err != nil {
		if err != model.OrderDataNotFound {
			errMsg = err.Error()
			log.Fatal(err.Error())
			return
		} else {
			errMsg = err.Error()
			log.Fatal(errMsg)
			return
		}
	}

	fmt.Println(items)
//TODO:格式化成msg--------------------
	prtResp, err := prt.PrintString(msg)
	if err != nil {
		errMsg = err.Error()
		log.Fatal(errMsg)
		return
	}
	log.Println(prtResp)
	result, err = json.Marshal(prtResp)
	if err != nil {
		errMsg = err.Error()
		log.Fatal(errMsg)
		return
	}

}
