package main

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {
	Router = gin.Default()
	api := Router.Group("/api/v1")
	{
		api.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "test successful",
			})
		})
	}
	Router.Run(":5000")
}
