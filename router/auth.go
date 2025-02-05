package router

import (
	"go-laris/controllers"

	"github.com/gin-gonic/gin"
)

func Auth(rg *gin.RouterGroup) {
	rg.POST("/login", controllers.AuthLogin)
	rg.POST("/register", controllers.CreateUsers)

}
