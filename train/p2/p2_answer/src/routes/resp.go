package routes

import "time"

type MenuList struct {
	ShopName   string
	TableName  string
	RecId      uint64
	EatType    string
	FoodsList  []FoodList
	Memo       string
	TotalPrice float64
	CreateTime time.Time
}

type FoodList struct {
	Id        uint64
	Name      string
	UintPrice float64 //菜的单价
	Num       int
	Amount    float64 // 每种菜的总价
}

type PrtList struct {
	Id     int64
	PrtNo  string
	Secret string
}
