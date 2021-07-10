package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomePage(c *gin.Context) {
	data := map[string]interface{}{
		"msg": "HomePage",
	}
	c.JSON(http.StatusOK, data)
}

func NoRouter(c *gin.Context) {
	data := map[string]interface{}{
		"msg": "404 Not Found",
	}
	c.JSON(http.StatusNotFound, data)
}

func Json(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Weikoi",
		"msg":  "hello, golang",
		"age":  18,
	})
}
