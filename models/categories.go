package models

type Categories struct {
	Id   int    `json:"id"`
	Name string `json:"name_categories"`
}

type CategoriesPagination struct {
	Id             int    `json:"id"`
	Image          string `json:"image"`
	NameProduct    string `json:"name_product"`
	Price          int    `json:"price"`
	Discount       int    `json:"discount"`
	NameCategories string `json:"name_categories"`
}
