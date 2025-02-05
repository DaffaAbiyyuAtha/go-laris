package router

import (
	"go-laris/middlewares"

	"github.com/gin-gonic/gin"
)

func User(rg *gin.RouterGroup) {
	rg.Use(middlewares.AuthMiddleware())

}
