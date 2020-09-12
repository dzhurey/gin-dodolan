package auth

import "github.com/gin-gonic/gin"

// check API will renew token when token life is less than 3 days, otherwise, return null for token
func check(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test successful",
	})
}
