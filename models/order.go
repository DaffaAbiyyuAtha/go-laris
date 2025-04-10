package models

import "time"

type Order struct {
	ID              int
	OrderID         string
	UserID          int
	TotalPrice      int
	PaymentStatus   string
	TransactionTime time.Time
	Items           []OrderItem
}

type OrderItem struct {
	ID        int       `json:"id"`
	OrderID   string    `json:"order_id"`
	ProductID int       `json:"product_id"`
	Qty       int       `json:"qty"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductItem struct {
	ProductID int `json:"product_id"`
	Qty       int `json:"qty"`
}

type CreateOrderRequest struct {
	Products []ProductItem `json:"products"`
}

type OrderResponse struct {
	OrderID         string              `json:"order_id"`
	TotalPrice      int                 `json:"total_price"`
	PaymentStatus   string              `json:"payment_status"`
	TransactionTime time.Time           `json:"transaction_time"`
	User            UserInfo            `json:"user"`
	OrderItems      []OrderItemResponse `json:"order_items"`
}

type UserInfo struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type OrderItemResponse struct {
	ProductID    int    `json:"product_id"`
	ProductImage string `json:"product_image"`
	Qty          int    `json:"qty"`
	Price        int    `json:"price"`
}
