package router

import "github.com/gin-gonic/gin"

func RouterCombain(r *gin.Engine) {
	CategoriesRouter(r.Group("/categories"))
	AuthUser(r.Group("/login"))

}
