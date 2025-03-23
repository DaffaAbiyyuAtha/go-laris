package router

import (
	"go-laris/controllers"

	"github.com/gin-gonic/gin"
)

func CategoriesRouter(rg *gin.RouterGroup) {
	rg.GET("", controllers.CategoriesController)
	rg.GET("/filter", controllers.ListProductCategory)
}
