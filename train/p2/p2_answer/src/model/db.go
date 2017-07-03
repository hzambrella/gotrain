package model

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type OrderDB interface {
	Close()
	GetItemsByShopId(shopId uint64) ([]Item, error)
	GetRecordById(recordId uint64) (*Record, error)
	GetTableById(tableId string) (*Table, error)
	GetShopById(shopId uint64) (*Shop, error)
	GetPrtsByShopId(shopId uint64) ([]Prt, error)
}

var (
	driver       string = "mysql"
	DataNotFound        = errors.New("data not found")
)

type City struct {
	Code        string
	Name        string
	FirstLetter string
}

type orderDB struct {
	*sql.DB
}

func NewOrderDB(dsn string) (OrderDB, error) {

	/*
		defer func(){
			if err:=recover();err!=nil{
				fmt.Println(err)
			}
		}()
	*/

	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	orderdb := &orderDB{db}
	return orderdb, nil
}

func (db *orderDB) Close() {
	db.Close()
}

const (
	getItemByShopIdSql = `
	SELECT
		id,name,sort,shop_id
	FROM
		order_item
	WHERE
		shop_id=?
	ORDER BY
		sort

	`
)

func (db *orderDB) GetItemsByShopId(shopId uint64) ([]Item, error) {
	rows, err := db.Query(getItemByShopIdSql, shopId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []Item{}
	for rows.Next() {
		result := Item{}
		if err := rows.Scan(
			&result.Id,
			&result.Name,
			&result.Sort,
			&result.ShopId,
		); err != nil {
			return nil, err
		}
		result.ShopId = shopId
		results = append(results, result)
	}
	if len(results) == 0 {
		return nil, DataNotFound
	}
	return results, nil

}

const (
	getRecordByIdSql = `
SELECT
	tb1.id, tb1.shop_id, tb1.price, tb1.eat_type, 
	tb1.pay_method, tb1.table_id, tb1.user_id, tb1.status,
	tb1.create_time, tb1.update_time, tb1.memo, 
	tb2.good_id,tb2.good_num,tb2.good_history_name,tb2.good_history_price
FROM
	order_record tb1
LEFT JOIN 
	order_record_detail tb2 
ON
	tb1.id=tb2.record_id
WHERE
	tb1.id=?
`
)

func (db *orderDB) GetRecordById(recordId uint64) (*Record, error) {
	rows, err := db.Query(getRecordByIdSql, recordId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := Record{}
	var goodId uint64
	var goodNum int
	var goodHisName string
	var goodHisPrice float64
	//	var time1 time.Time
	//	var time2 time.Time
	for rows.Next() {
		if err := rows.Scan(
			&result.Id,
			&result.ShopId,
			&result.Price,
			&result.EatType,
			&result.PayMethod,
			&result.TableId,
			&result.UserId,
			&result.Status,
			&result.CreateTime,
			&result.UpdateTime,
			//&time1,
			//&time2,
			&result.Memo,
			&goodId,
			&goodNum,
			&goodHisName,
			&goodHisPrice,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, DataNotFound
			}
			return nil, err
		}

		//		result.CreateTime = time1.Format("01-02 15:04")
		//		result.UpdateTime = time2.Format("01-02 15:04")
		result.GoodId = append(result.GoodId, goodId)
		result.GoodNum = append(result.GoodNum, goodNum)
		result.GoodHisName = append(result.GoodHisName, goodHisName)
		result.GoodHisPrice = append(result.GoodHisPrice, goodHisPrice)

	}
	return &result, nil
}

const (
	getTableByIdSql = `
	SELECT
		id, shop_id, table_name, qr_code 
	FROM
		order_table 
	WHERE
		id=?
`
)

func (db *orderDB) GetTableById(tableId string) (*Table, error) {
	var table Table
	if err := db.QueryRow(getTableByIdSql, tableId).Scan(
		&table.Id,
		&table.ShopId,
		&table.Name,
		&table.QrCode,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, DataNotFound
		}
		return nil, err
	}
	return &table, nil
}

const (
	getShopByIdSql = `
	SELECT
		id, name, message, ndagent, 
		nd_cash_store_id, status, create_time, update_time,logo
	FROM
		order_shop 
	WHERE
		id=?
`
)

func (db *orderDB) GetShopById(shopId uint64) (*Shop, error) {
	var shop Shop
	//var time1 time.Time
	//var time2 time.Time
	if err := db.QueryRow(getShopByIdSql, shopId).Scan(
		&shop.Id,
		&shop.Name,
		&shop.Message,
		&shop.NdAgent,
		&shop.NdcsId,
		&shop.Status,
		&shop.CreateTime,
		&shop.UpdateTime,
		//&time1,
		//&time2,
		&shop.Logo,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, DataNotFound
		}
		return nil, err
	}
	//shop.CreateTime = time1.Format("01-02 15:04")
	//shop.UpdateTime = time2.Format("01-02 15:04")
	return &shop, nil
}

const (
	getPrtsByShopIdSql = `
	SELECT
		id,prt_no,shop_id,secret,status,create_time,update_time,memo
	FROM
		order_prt
	WHERE
		shop_id=?
	AND
		status=1

 `
)

func (db *orderDB) GetPrtsByShopId(shopId uint64) ([]Prt, error) {
	rows, err := db.Query(getPrtsByShopIdSql, shopId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	//	var time1 time.Time
	//	var time2 time.Time
	results := []Prt{}
	for rows.Next() {
		result := Prt{}
		if err := rows.Scan(
			&result.Id,
			&result.PrtNo,
			&result.ShopId,
			&result.Secret,
			&result.Status,
			&result.CreateTime,
			&result.UpdateTime,
			//	&time1,
			//	&time2,
			&result.Memo,
		); err != nil {
			return nil, err
		}
		//		result.CreateTime = time1.Format("01-02 15:04")
		//		result.UpdateTime = time2.Format("01-02 15:04")
		results = append(results, result)
	}
	if len(results) == 0 {
		return nil, err
	}
	return results, nil
}
