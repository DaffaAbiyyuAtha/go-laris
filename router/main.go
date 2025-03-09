package router

import "github.com/gin-gonic/gin"

func RouterCombain(r *gin.Engine) {
	CategoriesRouter(r.Group("/categories"))
	Auth(r.Group("/auth"))
	User(r.Group("/user"))
	Wishlist(r.Group("/wishlist"))
	Product(r.Group("/product"))
}
