package model

type Item struct {
	Id     uint64 `json:"id"`
	Name   string `json:"name"`
	Sort   int    `json:"sort"`
	ShopId uint64 `json:"shopId"`
}
