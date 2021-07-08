package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.FormValue("token")
		if token == "" {
			c.JSON(401, gin.H{
				"message": "Token required",
			})
			c.Abort()
			return
		}
		if token != "accesstoken" {
			c.JSON(http.StatusOK, gin.H{
				"message": "Invalid Token",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
