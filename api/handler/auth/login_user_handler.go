package handler

import (
	"fmt"
	"forum-api/api/handler"
	request "forum-api/api/request/auth"
	userModel "forum-api/internal/model/user"
	"forum-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func LoginUserHandler(ctx *gin.Context) {
	userRequest := request.LoginUserRequest{}

	err := ctx.ShouldBindJSON(&userRequest)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": handler.ValidateError(err),
		})
		return
	}

	user, err := userModel.FindByEmail(userRequest.Email)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}

	if !utils.ComparePasswords(user.Password, []byte(userRequest.Password)) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Password dont match",
		})
		return
	}

	secret := []byte(os.Getenv("SECRET"))

	token, err := utils.CreateJWT(secret, user.ID)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao gerar token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": token,
	})
}
