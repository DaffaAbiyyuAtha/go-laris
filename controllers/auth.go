package controllers

import (
	"fmt"
	"go-laris/dtos"
	"go-laris/lib"
	"go-laris/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthLogin(ctx *gin.Context) {
	var user dtos.User
	ctx.Bind(&user)

	if len(user.Password) < 8 {
		ctx.JSON(http.StatusBadRequest, lib.Respont{
			Success: false,
			Message: "Password must be at least 8 characters long",
			Result:  nil,
		})
		return
	}

	found := repository.FindOneUserByEmail(user.Email)
	fmt.Println(found, "sini")

	if found == (dtos.User{}) {
		ctx.JSON(http.StatusUnauthorized, lib.Respont{
			Success: false,
			Message: "Wrong Email",
			Result:  nil,
		})
		return
	}

	isVerified := lib.Verify(user.Password, found.Password)

	if isVerified {
		JWT := lib.GenerateUserTokenById(found.Id)
		ctx.JSON(http.StatusOK, lib.Respont{
			Success: true,
			Message: "Login Success",
			Result:  dtos.Token{Token: JWT},
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, lib.Respont{
			Success: false,
			Message: "Wrong Password",
			Result:  nil,
		})
	}
}
