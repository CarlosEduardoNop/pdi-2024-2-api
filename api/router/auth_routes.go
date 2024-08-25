package router

import (
	auth "forum-api/api/handler/auth"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(routes *gin.RouterGroup) {
	orderRoutes := routes.Group("/auth")
	{
		orderRoutes.POST("/register", auth.RegisterUserHandler)
		orderRoutes.POST("/login", auth.LoginUserHandler)
	}
}
