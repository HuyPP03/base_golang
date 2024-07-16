package controllers

import (
	"net/http"
	"strings"

	"github.com/HuyPP03/learn/src/interfaces"
	"github.com/HuyPP03/learn/src/services"
	"github.com/HuyPP03/learn/src/utils"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerData interfaces.RegisterRequest
	if err := c.ShouldBindJSON(&registerData); err != nil {
		error := strings.Split(err.Error(), "Error:")
		c.JSON(http.StatusBadRequest, utils.NewErrorResponse(error[1], http.StatusBadRequest))
		return
	}

	user, err := services.Register(registerData.Username, registerData.Email, registerData.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse(user, "Register successfully!", http.StatusCreated))
}

func Login(c *gin.Context) {
	var loginData interfaces.LoginRequest
	if err := c.ShouldBindJSON(&loginData); err != nil {
		error := strings.Split(err.Error(), "Error:")
		c.JSON(http.StatusBadRequest, utils.NewErrorResponse(error[1], http.StatusBadRequest))
		return
	}

	token, err := services.Login(loginData.Email, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.NewErrorResponse(err.Error(), http.StatusUnauthorized))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse(token, "Success", http.StatusOK))
}

func GetProfile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	claims := user.(*utils.Claims) // Type assertion

	c.JSON(http.StatusOK, utils.NewSuccessResponse(claims.UserID, "Success", http.StatusOK))
}

func Upload(c *gin.Context) {
	filename, err := utils.UploadFile(c, "file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "filename": filename})
}

func Uploads(c *gin.Context) {
	uploadedFiles, err := utils.UploadFiles(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Files uploaded successfully", "uploadedFiles": uploadedFiles})
}
