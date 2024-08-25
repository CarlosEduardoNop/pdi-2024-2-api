package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	InitializeRoutes(router)

	router.Run(":" + os.Getenv("APP_PORT"))
}
