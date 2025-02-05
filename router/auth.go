package router

import (
	"go-laris/controllers"

	"github.com/gin-gonic/gin"
)

func AuthUser(rg *gin.RouterGroup) {
	rg.POST("", controllers.User)
}
