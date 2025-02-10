package controllers

import (
	"go-laris/lib"
	"go-laris/models"
	"go-laris/repository"
	"log"
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

func DeleteWishlist(ctx *gin.Context) {
	userId := ctx.GetInt("userId")

	productId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		lib.HandlerBadReq(ctx, "Invalid Product Id")
		return
	}

	err = repository.DeleteWishlist(userId, productId)
	if err != nil {
		if err.Error() == "wishlist item not found" {
			lib.HandlerNotfound(ctx, "Wishlist Item Not Found")
			return
		}
		lib.HandlerBadReq(ctx, "Failed To Detele Wishlist Item")

		return
	}
	lib.HandlerOK(ctx, "Wishlist Item Deleted Successfully", nil, nil)

}
