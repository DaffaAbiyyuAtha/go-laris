package models

type Product struct {
	Id             int      `json:"id"`
	NameProduct    string   `json:"name_product"`
	Price          int      `json:"price"`
	Discount       int      `json:"discount"`
	Description    string   `json:"description"`
	CategorId      int      `json:"categories_id"`
	NameCategories string   `json:"name_categories"`
	Images         []string `json:"images"`
}
