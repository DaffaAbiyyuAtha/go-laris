package router

import (
	"go-laris/controllers"
	"go-laris/middlewares"

	"github.com/gin-gonic/gin"
)

func User(rg *gin.RouterGroup) {
	rg.GET("/owner/manage", controllers.GetUsersfoOwner)
	rg.GET("/owner/manage/search", controllers.GetUsersByFullNamefoOwner)
	rg.DELETE("/owner/manage/delete/:id", controllers.DeleteUserforOwner)
	rg.GET("/admin/manage", controllers.GetUsersfoAdmin)
	rg.GET("/admin/manage/search", controllers.GetUsersByFullNamefoAdmin)
	rg.Use(middlewares.AuthMiddleware())
	rg.GET("", controllers.FindAllUser)
	rg.PATCH("/update", controllers.UpdateProfile)
}
