package router

import (
	"forum-api/api/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	routes := router.Group("/api/" + os.Getenv("API_VERSION"))

	routes.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	RegisterAuthRoutes(routes)

	authenticatedRoutes := routes.Group("", middleware.ValidateAuthorizationMiddleware)
	{
		RegisterForumRoutes(authenticatedRoutes)
	}
}
