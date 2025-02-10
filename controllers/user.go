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

func CreateUsers(ctx *gin.Context) {
	account := dtos.JoinRegist{}
	if err := ctx.ShouldBind(&account); err != nil {
		lib.HandlerBadReq(ctx, err.Error())
		return
	}

	if _, err := govalidator.ValidateStruct(account); err != nil {
		lib.HandlerBadReq(ctx, "Validation error: "+err.Error())
		return
	}

	profile, err := repository.CreateUser(account)
	if *account.Email == "" && account.Password == "" && profile.FullName == "" {
		lib.HandlerBadReq(ctx, "DataerHandlerBadReq")
		return
	}

	if err != nil {
		lib.HandlerBadReq(ctx, err.Error())
		return
	}

	lib.HandlerOK(ctx, "Register User success", nil, gin.H{
		"id":       profile.Id,
		"fullname": profile.FullName,
		"email":    account.Email,
		"gender":   profile.Gender,
		"role_id":  account.RoleId,
	})
}

func FindAllUser(ctx *gin.Context) {
	user := repository.FindAllUser()

	lib.HandlerOK(ctx, "List All User", user, nil)

}

func UpdateProfile(ctx *gin.Context) {

	id := ctx.GetInt("userId")
	if id == 0 {
		lib.HandlerBadReq(ctx, "User ID not found")
		return
	}

	var form dtos.Profile
	if err := ctx.ShouldBind(&form); err != nil {
		lib.HandlerBadReq(ctx, "Invalid input data")
		return
	}

	file, err := ctx.FormFile("image")
	var img *string

	if err == nil {
		allowExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
		fileExt := strings.ToLower(filepath.Ext(file.Filename))

		if !allowExt[fileExt] {
			lib.HandlerBadReq(ctx, "Invalid file extension")
			return
		}

		image := uuid.New().String() + fileExt
		root := "./img/profile"

		if err := ctx.SaveUploadedFile(file, root+image); err != nil {
			lib.HandlerBadReq(ctx, "Upload image failed")
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
		lib.HandlerBadReq(ctx, "Update profile failed")
		return
	}

	profile, err := repository.FindOneProfile(id)
	if err != nil {
		lib.HandlerBadReq(ctx, "Failed to find profile")
		return
	}

	userData := repository.FindOneUser(id)

	lib.HandlerOK(ctx, "Profile updated successfully", nil, gin.H{
		"profile": profile,
		"user":    userData,
	})
}
