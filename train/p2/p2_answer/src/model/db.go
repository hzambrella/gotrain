package model

import (
	"errors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


var (
	driver string = "mysql"
	OrderDataNotFound=errors.New("data not found")
)

const (
	getCitysByFirstLetterSql = `
	SELECT
		code,name,first_letter
	FROM
		life_city
	WHERE
		first_letter=?
	`
)

type City struct {
	Code        string
	Name        string
	FirstLetter string
}

type OrderDB  struct{
	*sql.DB
}

func NewOrderDB(dsn string) (*OrderDB, error) {

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
	orderdb:=&OrderDB{db}
	return orderdb, nil
}

func (db *OrderDB) Close() {
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

func (db *OrderDB) GetItemsByShopId(shopId uint64) ([]Item, error) {
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
		return nil, OrderDataNotFound
	}
	return results, nil

}
