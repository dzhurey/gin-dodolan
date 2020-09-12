package main

import (
	"os"

	"gin-dodolan/api"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load .env environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("APP_PORT")
	router := gin.Default()
	api.ApplyRoutes(router)
	router.Run(":" + port)
}
