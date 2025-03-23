package controllers

import (
	"go-laris/lib"
	"go-laris/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CategoriesController(ctx *gin.Context) {
	categories := repository.FindAllCategories()

	lib.HandlerOK(ctx, "List All Categories", categories, nil)

}

func ListProductCategory(c *gin.Context) {
	search := c.Query("search")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 8
	}
	products, err := repository.GetFilterProductWithCategory(search, page, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Respont{
			Success: false,
			Message: "Failed to find Products",
		})
		return
	}
	log.Println("Produk yang diambil:", products)
	c.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "List Products",
		Result:  products,
	})
}
