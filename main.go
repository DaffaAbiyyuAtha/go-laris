package main

import (
	"go-laris/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/img/profile", "./img/profile")

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{
		"Origin", "Content-Type", "Authorization", "Content-Length",
	}
	r.Use(cors.New(corsConfig))
	router.RouterCombain(r)
	r.Run(":8080")
}
