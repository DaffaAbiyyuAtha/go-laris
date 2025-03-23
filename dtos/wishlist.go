package dtos

type Wishlist struct {
	Id           int      `json:"id"`
	ProfileId    *int     `json:"profile_id" db:"profile_id" form:"profile_id"`
	ProductId    *int     `json:"product_id" db:"product_id" form:"product_id"`
	NameProduct  *string  `json:"name_product" db:"name_product" form:"name_product"`
	Image        []string `json:"image" db:"image" form:"image"`
	NameCategory *string  `json:"name_categories" db:"name_categories" form:"name_categories"`
}
