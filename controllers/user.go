package controllers

import (
	"fmt"
	"go-laris/dtos"
	"go-laris/lib"
	"go-laris/repository"
	"path/filepath"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		lib.HandlerBadReq(c, "DataerHandlerBadReq")
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

func UpdateProfile(c *gin.Context) {
	id := c.GetInt("userId")
	var form dtos.Profile
	var user dtos.User

	err := c.Bind(&form)
	errUser := c.Bind(&user)
	data, _ := repository.FindOneProfile(id)
	dataProfile := repository.FindOneUser(id)

	if err != nil {
		lib.HandlerBadReq(c, "Invalid input data")
		return
	}

	if errUser != nil {
		lib.HandlerBadReq(c, "Failed user")
		return
	}

	lib.HandlerOK(c, "Profile Found", nil, gin.H{
		"profile": data,
		"user":    dataProfile,
	})
}

func UploadProfileImage(c *gin.Context) {
	id := c.GetInt("userId")
	fmt.Println(id)

	file, err := c.FormFile("image")
	if err != nil {
		lib.HandlerBadReq(c, "no files uploaded")
		return
	}

	allowExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
	fileExt := strings.ToLower(filepath.Ext(file.Filename))
	if !allowExt[fileExt] {
		lib.HandlerBadReq(c, "invalid file extension")
		return
	}

	image := uuid.New().String() + fileExt

	root := "./img/profile/"

	if err := c.SaveUploadedFile(file, root+image); err != nil {
		fmt.Println(err)

		lib.HandlerBadReq(c, "Upload image failed")

		return
	}

	img := "http://localhost:8080/img/profile/" + image
	result, err := repository.UpdateProfileImage(dtos.Profile{Picture: &img}, id)

	if err != nil {
		lib.HandlerBadReq(c, "Update image failed")
		return
	}

	lib.HandlerOK(c, "Upload image success", nil, result)
}
