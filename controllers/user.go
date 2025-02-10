package controllers

import (
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
		"gender":   profile.Gender,
		"role_id":  account.RoleId,
	})
}

func FindAllUser(c *gin.Context) {
	user := repository.FindAllUser()

	lib.HandlerOK(c, "List All User", user, nil)

}

func UpdateProfile(c *gin.Context) {

	id := c.GetInt("userId")
	if id == 0 {
		lib.HandlerBadReq(c, "User ID not found")
		return
	}

	var form dtos.Profile
	if err := c.ShouldBind(&form); err != nil {
		lib.HandlerBadReq(c, "Invalid input data")
		return
	}

	file, err := c.FormFile("image")
	var img *string

	if err == nil {
		allowExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
		fileExt := strings.ToLower(filepath.Ext(file.Filename))

		if !allowExt[fileExt] {
			lib.HandlerBadReq(c, "Invalid file extension")
			return
		}

		image := uuid.New().String() + fileExt
		root := "../img/profile"

		if err := c.SaveUploadedFile(file, root+image); err != nil {
			lib.HandlerBadReq(c, "Upload image failed")
			return
		}

		baseURL := "http://localhost:8080"
		imgUrl := baseURL + "/img/profile/" + image
		img = &imgUrl
	}

	profileData := dtos.Profile{
		Picture:    img,
		FullName:   form.FullName,
		Province:   form.Province,
		City:       form.City,
		PostalCode: form.PostalCode,
		Gender:     form.Gender,
		Country:    form.Country,
		Mobile:     form.Mobile,
		Address:    form.Address,
		UserId:     id,
	}

	_, err = repository.UpdateProfile(profileData, id)
	if err != nil {
		lib.HandlerBadReq(c, "Update profile failed")
		return
	}

	profile, err := repository.FindOneProfile(id)
	if err != nil {
		lib.HandlerBadReq(c, "Failed to find profile")
		return
	}

	userData := repository.FindOneUser(id)

	lib.HandlerOK(c, "Profile updated successfully", nil, gin.H{
		"profile": profile,
		"user":    userData,
	})
}
