package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 默认json路由
	r.GET("/json", func(context *gin.Context) {
		data := map[string]interface{}{
			"name": "Weikoi",
			"msg":  "hello, golang",
			"age":  18,
		}

		context.JSON(http.StatusOK, data)
	})

	// homePage
	r.GET("", func(context *gin.Context) {

		data := map[string]interface{}{
			"name": "Home",
		}
		context.JSON(http.StatusOK, data)
	})

	// 无效路由
	r.NoRoute(func(context *gin.Context) {
		data := map[string]interface{}{
			"name": "NotFound",
		}
		context.JSON(http.StatusNotFound, data)
	})

	// 路由组
	bookGroup := r.Group("/book")
	{
		bookGroup.GET("/addBook", func(context *gin.Context) {

		})
		bookGroup.GET("/getBook", func(context *gin.Context) {

		})
		bookGroup.GET("/updateBook", func(context *gin.Context) {

		})
		bookGroup.GET("/deleteBook", func(context *gin.Context) {

		})
	}

	r.Run()
}
