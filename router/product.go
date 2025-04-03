package router

import (
	"go-laris/controllers"
	"go-laris/middlewares"

	"github.com/gin-gonic/gin"
)

func Product(rg *gin.RouterGroup) {
	rg.GET("/", controllers.FindAllProduct)
	rg.GET("/page", controllers.ListProductName)
	rg.GET("/filter", controllers.ListAllFilterProduct)
	rg.GET("/home", controllers.ListProduct)
	rg.GET("/:id", controllers.FindProduct)
	rg.Use(middlewares.AuthMiddleware())
	// rg.POST("/", controllers.CreateProduct)
	rg.DELETE("/:id", controllers.DeleteProduct)
}
