package router

import (
	"go-laris/controllers"
	"go-laris/middlewares"

	"github.com/gin-gonic/gin"
)

func User(rg *gin.RouterGroup) {
	rg.Use(middlewares.AuthMiddleware())
	rg.GET("", controllers.FindAllUser)
	rg.PATCH("/update", controllers.UpdateProfile)
	rg.PATCH("/img", controllers.UploadProfileImage)

}
