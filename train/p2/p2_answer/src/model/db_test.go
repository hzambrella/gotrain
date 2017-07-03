package model

import (
	"fmt"
	"testing"
	"engine/datastore"
	"log"
)

var db *OrderDB

func init(){
	d, err := datastore.ParseDataFromFile("../key.cfg")
	if err != nil {
		log.Fatal(err)
	}
	dsn:=d["dsn"]
	db,err=NewOrderDB(dsn)
	if err!=nil{
		fmt.Println(err.Error())
	}
}

func TestGetItemsByShopId(t *testing.T) {
	items, err := db.GetItemsByShopId(1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(items)
}
