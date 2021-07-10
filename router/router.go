package router

import (
	"GinProject/controller"
	"GinProject/controller/book"
	"GinProject/controller/user"
	"GinProject/middleware"
	"github.com/gin-gonic/gin"
)

func Init() {
	// Creates a default gin router
	r := gin.Default()

	// 全局中间件
	r.Use(middleware.TimerMiddleware())
	r.Use(middleware.ValidateToken())

	// 默认json路由(测试)
	r.GET("/json", controller.Json)

	// HomePage route
	// 主页路由
	r.GET("", controller.HomePage)

	// No route
	// 无效路由
	r.NoRoute(controller.NoRouter)

	// Grouping routes
	// 用户路由组
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", user.Register)
		userGroup.POST("/login", user.Login)
		userGroup.POST("/update", user.UpdateInfo)
		userGroup.POST("/logout", user.Logout)
		userGroup.POST("/delete", user.Delete)
	}

	// 图书路由组
	bookGroup := r.Group("/book")
	{
		bookGroup.POST("/add", Book.Add)
		bookGroup.GET("/get", Book.Get)
		bookGroup.POST("/update", Book.Update)
		bookGroup.POST("/delete", Book.Delete)
	}

	// restful api group：v1
	v1 := r.Group("/v1")
	{
		v1.GET("/home", controller.HomePage)
	}

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
