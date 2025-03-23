package controllers

import (
	"fmt"
	"go-laris/lib"
	"go-laris/models"
	"go-laris/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindAllWishlist(ctx *gin.Context) {
	results := repository.FindAllWishlist()
	lib.HandlerOK(ctx, "Find All Wishilist", nil, results)
}

func FindOneWishlist(ctx *gin.Context) {
	id := ctx.GetInt("userId")

	wishlist, err := repository.FindOneWishlist(id)
	if err != nil {
		lib.HandlerNotfound(ctx, "Wishhlist Not Found")
	}

	if len(wishlist) == 0 {
		lib.HandlerNotfound(ctx, "No Wishlist Found For This User")
	}

	var results []gin.H

	for _, wishlist := range wishlist {
		product, err := repository.FindOneProductById(wishlist.Id)
		if err != nil {
			log.Printf("Failed to fetch event with id %d: %v", wishlist.ProductId, err)
			continue
		}

		results = append(results, gin.H{
			"wishlist": wishlist,
			"product":  product,
		})
	}
	lib.HandlerOK(ctx, "Wishlist and events found", nil, results)
}

func CreateWishlist(ctx *gin.Context) {
	var newWishlist models.Wishlist

	id, exists := ctx.Get("userId")
	if !exists {
		lib.HandlerUnauthorized(ctx, "Unauthorized")
	}

	userId, ok := id.(int)
	if !ok {
		lib.HandlerStatusInternalServerError(ctx, "Invalid User Id")
	}

	productId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		lib.HandlerBadReq(ctx, "Invalid Product Id")
	}

	err = repository.CreateWishlist(productId, userId)
	if err != nil {
		log.Printf("Create Wishlist Error: %v", err)
		if err.Error() == "Wishlist Entry Already Exists" {
			lib.HandlerStatusConflict(ctx, "Product Is Already In Your Wishlist")
		}

		lib.HandlerBadReq(ctx, "Failed to create Wishlist")
		return
	}

	newWishlist.UserId = userId
	newWishlist.ProductId = productId

	lib.HandlerOK(ctx, "Wishlist Create Successfully", nil, newWishlist)

}

// func DeleteWishlist(ctx *gin.Context) {
// 	userId := ctx.GetInt("userId")

// 	productId, err := strconv.Atoi(ctx.Param("id"))
// 	if err != nil {
// 		lib.HandlerBadReq(ctx, "Invalid Product Id")
// 		return
// 	}

// 	err = repository.DeleteWishlist(userId, productId)
// 	if err != nil {
// 		if err.Error() == "wishlist item not found" {
// 			lib.HandlerNotfound(ctx, "Wishlist Item Not Found")
// 			return
// 		}
// 		lib.HandlerBadReq(ctx, "Failed To Detele Wishlist Item")

// 		return
// 	}
// 	lib.HandlerOK(ctx, "Wishlist Item Deleted Successfully", nil, nil)

// }

func FindWishlistbyProfileId(ctx *gin.Context) {
	profileId := ctx.GetInt("userId")

	dataWishlist, err := repository.FindWishlistByProfileId(profileId)
	if err != nil {
		fmt.Println("Error:", err)
		ctx.JSON(http.StatusNotFound, lib.Respont{
			Success: false,
			Message: "Wishlist Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "Wishlist Found",
		Result:  dataWishlist,
	})
}

func CreateNewWishlist(ctx *gin.Context) {
	profileId := ctx.GetInt("userId")

	productIdStr := ctx.PostForm("product_id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Respont{
			Success: false,
			Message: "Invalid product_id",
		})
		return
	}

	err = repository.CreateWishlist(profileId, productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Respont{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, lib.Respont{
		Success: true,
		Message: "Wishlist created successfully",
	})
}

func DeleteWishlist(ctx *gin.Context) {
	profileId := ctx.GetInt("userId")
	fmt.Println("Profile ID:", profileId)

	productID := ctx.Query("product_id")
	fmt.Println("Product ID dari query:", productID)
	if productID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "product_id is required",
		})
		return
	}

	productId, err := strconv.Atoi(productID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Respont{
			Success: false,
			Message: "Invalid product_id",
		})
		return
	}

	fmt.Println("Product ID yang mau dihapus:", productId)
	err = repository.DeleteWishlist(profileId, productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Respont{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "Wishlist deleted successfully",
	})
}

func GetWishlistByProfileAndProductName(ctx *gin.Context) {
	profileId := ctx.GetInt("userId")
	productName := ctx.Query("product_name")

	if productName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "product_name is required",
		})
		return
	}

	wishlist, err := repository.GetWishlistByProfileAndProductName(profileId, productName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Wishlist found",
		"result":  wishlist,
	})
}
