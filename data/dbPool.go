package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"

type UserInfo struct {
	UserId   int
	UserName string
	Password string
	Auth     int
}

func main() {
	db, err := gorm.Open("mysql", "root:admin@tcp(localhost:3306)/dev")
	var user_info UserInfo
	db.Raw("SELECT user_id, user_name, password, auth FROM users_info").Scan(&user_info)
	fmt.Println(user_info.Auth)
	if err != nil {
		panic(err)
	}

	defer db.Close()

}
