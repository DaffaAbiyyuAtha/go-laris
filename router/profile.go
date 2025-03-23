package router

import (
	"go-laris/controllers"
	"go-laris/middlewares"

	"github.com/gin-gonic/gin"
)

func ProfileRouter(rg *gin.RouterGroup) {
	rg.Use(middlewares.AuthMiddleware())
	rg.GET("", controllers.SeeOneProfileByUserId)
	rg.PATCH("/update", controllers.UpdateUserProfileController)
	rg.PATCH("/picture", controllers.UpdateProfilePicture)
}
