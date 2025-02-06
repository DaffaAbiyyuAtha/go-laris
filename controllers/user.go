package controllers

import (
	"go-laris/dtos"
	"go-laris/lib"
	"go-laris/repository"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func CreateUsers(c *gin.Context) {
	account := dtos.JoinRegist{}
	if err := c.ShouldBind(&account); err != nil {
		lib.HandlerBadReq(c, err.Error())
		return
	}

	if _, err := govalidator.ValidateStruct(account); err != nil {
		lib.HandlerBadReq(c, "Validation error: "+err.Error())
		return
	}
	profile, err := repository.CreateUser(account)
	if *account.Email == "" && account.Password == "" && profile.FullName == "" {
		lib.HandlerBadReq(c, "Data bad request")
		return
	}

	if err != nil {
		lib.HandlerBadReq(c, err.Error())
		return
	}

	lib.HandlerOK(c, "Register User success", nil, gin.H{
		"id":       profile.Id,
		"fullname": profile.FullName,
		"email":    account.Email,
		"role_id":  account.RoleId,
	})
}

func FindAllUser(c *gin.Context) {
	user := repository.FindAllUser()

	lib.HandlerOK(c, "List All User", user, nil)

}
