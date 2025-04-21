package controllers

import (
	"fmt"
	"go-laris/lib"
	"go-laris/models"
	"go-laris/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateOrder(ctx *gin.Context) {
	var req models.CreateOrderRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Respont{
			Success: false,
			Message: "Invalid request: " + err.Error(),
		})
		return
	}

	userID := ctx.GetInt("userId")

	profile, err := repository.FindOneProfile(userID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Respont{
			Success: false,
			Message: "Failed to find profile: " + err.Error(),
		})
		return
	}

	user := repository.FindOneUser(userID)

	var totalPrice int
	var orderItems []models.OrderItem

	for _, p := range req.Products {
		product, err := repository.FindOneProductById(p.ProductID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, lib.Respont{
				Success: false,
				Message: "Product not found: " + err.Error(),
			})
			return
		}

		price := product.Price
		if product.Discount > 0 {
			price = product.Price - (product.Price * product.Discount / 100)
		}

		totalPrice += price * p.Qty

		orderItems = append(orderItems, models.OrderItem{
			ProductID: p.ProductID,
			Qty:       p.Qty,
			Price:     price,
		})
	}

	orderID := uuid.New().String()
	now := time.Now()

	order := models.Order{
		OrderID:         orderID,
		UserID:          userID,
		TotalPrice:      totalPrice,
		PaymentStatus:   "pending",
		TransactionTime: now,
	}

	db := lib.DB()
	tx, err := db.Begin(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Respont{
			Success: false,
			Message: "Failed to start transaction: " + err.Error(),
		})
		return
	}

	defer tx.Rollback(ctx)

	if err := repository.CreateOrder(ctx, &order); err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Respont{
			Success: false,
			Message: "Failed to create order: " + err.Error(),
		})
		return
	}

	for i := range orderItems {
		orderItems[i].OrderID = orderID
	}

	for i := range orderItems {
		if err := repository.CreateOrderItem(ctx, &orderItems[i]); err != nil {
			ctx.JSON(http.StatusInternalServerError, lib.Respont{
				Success: false,
				Message: "Failed to create order item: " + err.Error(),
			})
			return
		}
	}

	if err := tx.Commit(ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Respont{
			Success: false,
			Message: "Failed to commit transaction: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "Order created successfully",
		Result: gin.H{
			"order_id":    order.OrderID,
			"total_price": order.TotalPrice,
			"customer": gin.H{
				"fullname": profile.FullName,
				"email":    user.Email,
			},
			"items": orderItems,
		},
	})
}

func GetAllOrders(c *gin.Context) {
	orders, err := repository.GetAllOrders(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.Respont{
			Success: false,
			Message: fmt.Sprintf("Failed to fetch orders: %v", err.Error()),
			Result:  nil,
		})
		return
	}

	c.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "Success See All Order",
		Result:  orders,
	})
}

func GetOrderByID(c *gin.Context) {
	orderID := c.Param("order_id")

	order, err := repository.FindOrderByID(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.Respont{
			Success: false,
			Message: fmt.Sprintf("Failed to fetch order: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "Order fetched successfully",
		Result:  order,
	})
}
