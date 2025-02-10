package router

import (
	"go-laris/controllers"
	"go-laris/middlewares"

	"github.com/gin-gonic/gin"
)

func Product(rg *gin.RouterGroup) {
	rg.GET("/", controllers.FindAllProduct)
	rg.Use(middlewares.AuthMiddleware())
	rg.POST("/", controllers.CreateProduct)
	rg.DELETE("/:id", controllers.DeleteProduct)
}
