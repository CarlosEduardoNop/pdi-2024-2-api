package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	UseCors(router)
	InitializeRoutes(router)

	router.Run(":" + os.Getenv("APP_PORT"))
}
