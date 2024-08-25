package handler

import (
	"fmt"
	"forum-api/api/handler"
	request "forum-api/api/request/auth"
	"forum-api/internal/model/user"
	"forum-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUserHandler(ctx *gin.Context) {
	userRequest := request.RegisterUserRequest{}

	err := ctx.ShouldBindJSON(&userRequest)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": handler.ValidateError(err),
		})
		return
	}

	hashPassword, err := utils.HashPassword(userRequest.Password)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": handler.ValidateError(err),
		})
		return
	}

	userCreated, err := user.Create(map[string]interface{}{
		"name":     userRequest.User,
		"email":    userRequest.Email,
		"password": string(hashPassword),
	})

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error to create user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": userCreated,
	})
}
