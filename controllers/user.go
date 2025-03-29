package controllers

import (
	"fmt"
	"go-laris/dtos"
	"go-laris/lib"
	"go-laris/repository"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUsers(ctx *gin.Context) {
	account := dtos.UserRegist{}
	if err := ctx.ShouldBind(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Respont{
			Success: false,
			Message: "Failed to bind data: " + err.Error(),
		})
		return
	}

	if len(account.Password) < 8 {
		ctx.JSON(http.StatusBadRequest, lib.Respont{
			Success: false,
			Message: "Password must be at least 8 characters long",
		})
		return
	}

	if _, err := govalidator.ValidateStruct(account); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Respont{
			Success: false,
			Message: lib.FormatValidationError(err.Error()),
		})
		return
	}

	if account.Email == account.Password {
		ctx.JSON(http.StatusBadRequest, lib.Respont{
			Success: false,
			Message: "Password cannot be the same as Email",
		})
		return
	}

	existingUser, _ := repository.FindOneUserByEmailForRegist(account.Email)
	if existingUser.Email != "" {
		ctx.JSON(http.StatusBadRequest, lib.Respont{
			Success: false,
			Message: "Email already registered",
		})
		return
	}

	joinAccount := dtos.JoinRegist{
		Email:    &account.Email,
		Password: account.Password,
		RoleId:   account.RoleId,
		Results: dtos.Profile{
			FullName: account.FullName,
		},
	}

	profile, err := repository.CreateUser(joinAccount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Respont{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "Register User success",
		Result: gin.H{
			"id":       profile.Id,
			"fullname": profile.FullName,
			"email":    account.Email,
			"role_id":  account.RoleId,
		},
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

func GetUsersfoOwner(ctx *gin.Context) {
	users, err := repository.FindUsersByRoleforOwner()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Respont{
			Success: false,
			Message: "Failed to fetch users: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "Find All Users successfully",
		Result:  users,
	})
}

func GetUsersByFullNamefoOwner(ctx *gin.Context) {
	fullName := ctx.Query("fullname")

	users, err := repository.FindManageUsersByFullName(fullName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Respont{
			Success: false,
			Message: "Failed to fetch users: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "Users fetched successfully",
		Result:  users,
	})
}

func DeleteUserforOwner(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Respont{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	profiles, err := repository.DeleteUserforOwner(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Respont{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "User deleted successfully",
		Result:  profiles,
	})
}

func GetUsersfoAdmin(ctx *gin.Context) {
	users, err := repository.FindUsersByRoleforAdmin()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Respont{
			Success: false,
			Message: "Failed to fetch users: " + err.Error(),
		})
		return
	}

	fmt.Println(users)

	ctx.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "Find All Users successfully",
		Result:  users,
	})
}

func GetUsersByFullNamefoAdmin(ctx *gin.Context) {
	fullName := ctx.Query("fullname")

	users, err := repository.FindManageUsersByFullNamefoAdmin(fullName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Respont{
			Success: false,
			Message: "Failed to fetch users: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "Users fetched successfully",
		Result:  users,
	})
}
