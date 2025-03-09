package dtos

type Product struct {
	Id           int      `json:"id"`
	NameProduct  string   `json:"nameProduct" form:"nameProduct" db:"name_product"`
	Price        int      `json:"price" form:"price" db:"price"`
	Discount     int      `json:"discount" form:"discount" db:"discount"`
	Description  string   `json:"description" form:"description" db:"description"`
	CategoriesId *int     `json:"categoriesId" form:"categoriesId" db:"categories_id"`
	Image        []string `json:"image" form:"image" db:"image"`
}
