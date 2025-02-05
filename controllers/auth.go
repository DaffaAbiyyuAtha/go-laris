package controllers

import (
	"go-laris/dtos"
	"go-laris/lib"
	"go-laris/repository"

	"github.com/gin-gonic/gin"
)

func User(ctx *gin.Context) {
	var user dtos.User
	ctx.Bind(&user)
	found := repository.FindOneUserByEmail(user.Email)

	if found == (dtos.User{}) {
		lib.HandlerUnauthorized(ctx, "Wrong Email")
		return
	}

	isVerified := lib.Verify(user.Password, found.Password)

	if isVerified {
		JWT := lib.GenerateUserTokenById(found.Id)
		lib.HandlerOK(ctx, "Login Success", nil, dtos.Token{Token: JWT})
	} else {
		lib.HandlerUnauthorized(ctx, "Wrong Password")

	}
}
