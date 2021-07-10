package checker

import (
	"GinProject/utils/exception"
	"github.com/gin-gonic/gin"
	"log"
)

func ParamsChecker(c *gin.Context, api string) {
	switch api {
	case "register":
		checkRegisterParams(c)

	}
}

func checkRegisterParams(c *gin.Context) {
	userName := c.PostForm("username")
	password := c.PostForm("password")
	repeatPassword := c.PostForm("repeat_password")

	log.Printf("用户名通过检验%v\n", userName)

	//if password!=repeatPassword{
	//	return errors.New("密码不一致")
	//}
	exception.Assert(password == repeatPassword, "两次密码必须一致")
}
