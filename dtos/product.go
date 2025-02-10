package dtos

type Product struct {
	Id           int     `json:"id"`
	Image        *string `json:"image" form:"image" db:"image"`
	NameProduct  string  `json:"nameProduct" form:"nameProduct" db:"name_product"`
	Price        int     `json:"price" form:"price" db:"price"`
	Discount     int     `json:"discount" form:"discount" db:"discount"`
	Total        int     `json:"total" form:"total" db:"total"`
	CategoriesId *int    `json:"categoriesId" form:"categoriesId" db:"categories_id"`
}
