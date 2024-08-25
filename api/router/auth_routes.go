package router

import (
	auth "forum-api/api/handler/auth"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(routes *gin.RouterGroup) {
	authRoutes := routes.Group("/auth")
	{
		authRoutes.POST("/register", auth.RegisterUserHandler)
		authRoutes.POST("/login", auth.LoginUserHandler)
	}
}
