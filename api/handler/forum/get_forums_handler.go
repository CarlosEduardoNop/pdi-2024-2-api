package handler

import (
	"forum-api/internal/model/forum"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetForumsHandler(ctx *gin.Context) {
	forum, err := forum.GetAllWithLastPost()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, forum)
}
