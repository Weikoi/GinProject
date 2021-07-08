package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func TimerMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		timeS := time.Now()
		context.Next()
		log.Printf("请求耗时:%v\n", time.Since(timeS))
	}
}
