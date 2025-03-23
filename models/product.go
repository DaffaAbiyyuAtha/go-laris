package models

type Product struct {
	Id             int      `json:"id"`
	NameProduct    string   `json:"nameProduct"`
	Price          int      `json:"price"`
	Discount       int      `json:"discount"`
	Description    string   `json:"description"`
	CategorId      int      `json:"categoriesId"`
	NameCategories string   `json:"name_categories"`
	Image          []string `json:"image"`
}
