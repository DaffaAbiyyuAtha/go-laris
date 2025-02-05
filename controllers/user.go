package controllers

import (
	"fmt"
	"go-laris/dtos"
	"go-laris/lib"
	"go-laris/repository"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func CreateUsers(c *gin.Context) {
	newUser := dtos.User{}

	if err := c.ShouldBind(&newUser); err != nil {
		lib.HandlerBadReq(c, "Invalid input data")
		return
	}

	if _, err := govalidator.ValidateStruct(newUser); err != nil {
		lib.HandlerBadReq(c, "Validation error: "+err.Error())
		return
	}
	addUser := repository.CreateUser(newUser)
	fmt.Println(addUser)
	lib.HandlerOK(c, "User created successfully", nil, addUser)
}
