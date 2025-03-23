package router

import (
	"go-laris/controllers"
	"go-laris/middlewares"

	"github.com/gin-gonic/gin"
)

func Wishlist(rg *gin.RouterGroup) {
	rg.Use(middlewares.AuthMiddleware())
	// rg.GET("", controllers.FindAllWishlist)
	rg.POST("/:id", controllers.CreateWishlist)
	// rg.GET("/:id", controllers.FindOneWishlist)
	rg.DELETE("/delete", controllers.DeleteWishlist)
	rg.POST("/create", controllers.CreateNewWishlist)
	rg.GET("/", controllers.FindWishlistbyProfileId)
	rg.GET("/search", controllers.GetWishlistByProfileAndProductName)
}
