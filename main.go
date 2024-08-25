package main

import (
	"forum-api/api/router"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router.Initialize()
}
