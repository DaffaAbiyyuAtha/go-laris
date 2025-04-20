package controllers

import (
	"fmt"
	"go-laris/dtos"
	"go-laris/lib"
	"go-laris/repository"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SeeOneProfileByUserId(ctx *gin.Context) {
	id := ctx.GetInt("userId")
	fmt.Println(id)
	dataProfile, err := repository.FindProfileByUserId(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Respont{
			Success: false,
			Message: "Failed to fetch profile",
		})
		return
	}
	fmt.Println("Profile:", dataProfile)
	dataUser, err := repository.FindUser(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Respont{
			Success: false,
			Message: "User Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "user Found",
		Result: gin.H{
			"profile": dataProfile,
			"user":    dataUser,
		},
	})
}

func UpdateUserProfileController(ctx *gin.Context) {
	userIDInterface, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "User ID not found",
		})
		return
	}
	userID := userIDInterface.(int)
	fmt.Println("User ID:", userID)

	var profile dtos.Profile

	if err := ctx.ShouldBind(&profile); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to bind data",
		})
		return
	}

	if profile.FullName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "You Must Fill Full Name",
		})
		return
	}

	updatedProfile, err := repository.UpdateUserProfile(userID, profile)
	if err != nil {
		fmt.Println("Update Error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to update profile",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Update Profile Successfully",
		"result":  updatedProfile,
	})
}

func UpdateProfilePicture(c *gin.Context) {
	id := c.GetInt("userId")

	oldProfile, err := repository.FindProfileByUserId(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, lib.Respont{
			Success: false,
			Message: "Failed to get old profile",
		})
		return
	}

	fullName := c.DefaultPostForm("fullname", "")
	province := c.DefaultPostForm("province", "")
	city := c.DefaultPostForm("city", "")
	postalCode := c.DefaultPostForm("postalCode", "")
	gender := c.DefaultPostForm("gender", "")
	country := c.DefaultPostForm("country", "")
	mobile := c.DefaultPostForm("mobile", "")
	address := c.DefaultPostForm("address", "")

	if oldProfile.Picture != nil {
		oldPicturePath := "./img/profile/" + filepath.Base(*oldProfile.Picture)
		if err := os.Remove(oldPicturePath); err != nil && !os.IsNotExist(err) {
			fmt.Println("Failed to delete old picture:", err)
		}
	}

	var picturePath string
	file, err := c.FormFile("picture")
	if err == nil {
		cek := map[string]bool{".jpg": true, ".png": true, ".jpeg": true}
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if !cek[ext] {
			c.JSON(http.StatusBadRequest, lib.Respont{
				Success: false,
				Message: "Invalid file format",
			})
			return
		}

		picture := uuid.New().String() + ext
		savePicture := "./img/profile/"

		if err := c.SaveUploadedFile(file, savePicture+picture); err != nil {
			c.JSON(http.StatusInternalServerError, lib.Respont{
				Success: false,
				Message: "Failed to save file",
			})
			return
		}

		picturePath = "http://localhost:8100/profile/picture/" + picture
	}

	updatedProfile, err := repository.UpdateProfilePicture(
		fullName, province, city, postalCode, gender, country, mobile, address, picturePath, id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.Respont{
			Success: false,
			Message: "Failed to update profile",
		})
		return
	}
	user, err := repository.FindUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.Respont{
			Success: false,
			Message: "Failed to get user",
		})
		return
	}

	c.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "Profile updated successfully",
		Result: gin.H{
			"profile": updatedProfile,
			"user":    user,
		},
	})
}
