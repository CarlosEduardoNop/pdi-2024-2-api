package router

import (
	forum "forum-api/api/handler/forum"
	"github.com/gin-gonic/gin"
)

func RegisterForumRoutes(routes *gin.RouterGroup) {
	forumRoutes := routes.Group("/forum")
	{
		forumRoutes.GET("", forum.GetForumsHandler)
	}
}
