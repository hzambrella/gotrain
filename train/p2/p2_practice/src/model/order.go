package model

import "time"

type Shop struct {
	Id         uint64    `json:"id"`
	Name       string    `json:"name"`
	Logo       string    `json:"logo"`
	Message    string    `json:"message"`
	NdAgent    string    `json:"ndAgent"`
	NdcsId     int       `json:"nd_cash_store_id"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

type Item struct {
	Id     uint64 `json:"id"`
	Name   string `json:"name"`
	Sort   int    `json:"sort"`
	ShopId uint64 `json:"shopId"`
}

type Good struct {
	Id          uint64    `json:"id"`
	ShopId      uint64    `json:"shopId"`
	ItemId      uint64    `json:"itemId"`
	Name        string    `json:"name"`
	PicSmall    string    `json:"picSmall"`
	PicBig      string    `json:"picBig"`
	Price       float64   `json:"price"`
	Discount    int       `json:"discount"`
	Sold        int       `json:"sold"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	Sort        int       `json:"sor"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
	Taste       string    `json:"taste"`
}

type Record struct {
	Id           uint64    `json:"id"`
	ShopId       uint64    `json:"shopId"`
	Price        float64   `json:"price"`
	EatType      int       `json:"eatType"`
	PayMethod    int       `json:"payMethod"`
	TableId      string    `json:"tableId"`
	UserId       string    `json:"userId"`
	Status       int       `json:"status"`
	CreateTime   time.Time `json:"createTime"`
	UpdateTime   time.Time `json:"updateTime"`
	Memo         string    `json:"memo"`
	GoodId       []uint64  `json:"goodId"`
	GoodNum      []int     `json:"goodNum"`
	GoodHisName  []string  `json:"goodName"`
	GoodHisPrice []float64 `json:"goodName"` // 单价(no discount)
}
type Records []Record

type Table struct {
	Id     string `json:"id"`
	ShopId uint64 `json:"shopId"`
	Name   string `json:"tableName"`
	QrCode string `json:"qrCode"`
}

type Prt struct {
	Id         int64     `json:"id"`
	PrtNo      string    `json:"prtno"`
	Secret     string    `json:"secret"`
	ShopId     int64     `json:"shopId"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	Memo       string    `json:"memo"`
}
