package repository

import (
	"context"
	"encoding/json"
	"go-laris/lib"
	"go-laris/models"
	"time"

	"github.com/jackc/pgx/v5"
)

func CreateOrder(ctx context.Context, order *models.Order) error {
	db := lib.DB()

	query := `
        INSERT INTO orders (order_id, user_id, total_price, payment_status, transaction_time)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
	err := db.QueryRow(ctx, query,
		order.OrderID,
		order.UserID,
		order.TotalPrice,
		order.PaymentStatus,
		order.TransactionTime,
	).Scan(&order.ID)

	return err
}

func CreateOrderItem(ctx context.Context, orderItem *models.OrderItem) error {
	db := lib.DB()

	query := `
		INSERT INTO order_items (order_id, product_id, qty, price, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at
	`

	now := time.Now()

	return db.QueryRow(ctx, query,
		orderItem.OrderID,
		orderItem.ProductID,
		orderItem.Qty,
		orderItem.Price,
		now,
		now,
	).Scan(&orderItem.ID, &orderItem.CreatedAt, &orderItem.UpdatedAt)
}

func GetAllOrders(ctx context.Context) ([]models.OrderResponse, error) {
	db := lib.DB()
	defer db.Close(ctx)

	query := `
		SELECT 
			o.order_id,
			o.total_price,
			o.payment_status,
			o.transaction_time,
			json_build_object(
				'fullname', p.fullname,
				'email', u.email
			) AS user,
			json_agg(
				json_build_object(
					'product_id', oi.product_id,
					'product_image', (
						SELECT pi.image 
						FROM product_images pi 
						WHERE pi.product_id = oi.product_id 
						ORDER BY pi.id ASC 
						LIMIT 1
					),
					'qty', oi.qty,
					'price', oi.price
				)
			) AS order_items
		FROM "orders" o
		JOIN "order_items" oi ON o.order_id = oi.order_id
		JOIN "product" pr ON oi.product_id = pr.id
		JOIN "user" u ON o.user_id = u.id
		JOIN "profile" p ON u.id = p.user_id
		GROUP BY o.order_id, o.total_price, o.payment_status, o.transaction_time, u.email, p.fullname
		ORDER BY o.transaction_time DESC

	`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.OrderResponse])
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func FindOrderByID(orderID string) (models.OrderResponse, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	var order models.OrderResponse
	var userJSON []byte
	var orderItemsJSON []byte

	sql := `
		SELECT 
			o.order_id,
			o.total_price,
			o.payment_status,
			o.transaction_time,
			json_build_object(
				'fullname', p.fullname,
				'email', u.email
			) AS user,
			json_agg(
				json_build_object(
					'product_id', oi.product_id,
					'product_image', (
						SELECT pi.image 
						FROM product_images pi 
						WHERE pi.product_id = oi.product_id 
						ORDER BY pi.id ASC 
						LIMIT 1
					),
					'qty', oi.qty,
					'price', oi.price
				)
			) AS order_items
		FROM "orders" o
		JOIN "order_items" oi ON o.order_id = oi.order_id
		JOIN "product" pr ON oi.product_id = pr.id
		JOIN "user" u ON o.user_id = u.id
		JOIN "profile" p ON u.id = p.user_id
		WHERE o.order_id = $1
		GROUP BY o.order_id, o.total_price, o.payment_status, o.transaction_time, u.email, p.fullname
	`

	row := db.QueryRow(context.Background(), sql, orderID)

	err := row.Scan(
		&order.OrderID,
		&order.TotalPrice,
		&order.PaymentStatus,
		&order.TransactionTime,
		&userJSON,
		&orderItemsJSON,
	)
	if err != nil {
		return order, err
	}

	_ = json.Unmarshal(userJSON, &order.User)
	_ = json.Unmarshal(orderItemsJSON, &order.OrderItems)

	return order, nil
}
