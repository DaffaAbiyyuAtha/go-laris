package models

type Wishlist struct {
	Id        int `json:"id"`
	UserId    int `json:"user_id" form:"user_id"`
	ProductId int `json:"product_id" form:"product_id"`
}
