package controllers

import (
	"go-laris/lib"
	"go-laris/repository"

	"github.com/gin-gonic/gin"
)

func CategoriesController(ctx *gin.Context) {
	categories := repository.FindAllCategories()

	lib.HandlerOK(ctx, "List All Categories", categories, nil)

}
