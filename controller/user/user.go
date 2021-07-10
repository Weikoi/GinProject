package user

import (
	"GinProject/controller/checker"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
}
func Logout(c *gin.Context) {
}

// 用户注册
func Register(c *gin.Context) {

	// params_checker
	checker.ParamsChecker(c, "register")

	userName := c.PostForm("username")
	password := c.PostForm("password")

}
func Delete(c *gin.Context) {
}
func UpdateInfo(c *gin.Context) {
}
