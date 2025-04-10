package router

import (
	"go-laris/controllers"
	"go-laris/middlewares"

	"github.com/gin-gonic/gin"
)

func OrderRouter(rg *gin.RouterGroup) {
	rg.GET("", controllers.GetAllOrders)
	rg.GET("/:order_id", controllers.GetOrderByID)
	rg.Use(middlewares.AuthMiddleware())
	rg.POST("", controllers.CreateOrder)
}
