package controllers

import (
	"go-laris/lib"
	"go-laris/repository"

	"github.com/gin-gonic/gin"
)

func CategoriesController(c *gin.Context) {
	categories := repository.FindAllCategories()

	lib.HandlerOK(c, "List All Categories", categories, nil)

}
