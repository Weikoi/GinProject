package data

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/mysql"

func main() {
	db, err := gorm.Open("mysql", "root:admin@admin(localhost:3306)/dev")

	if err != nil {
		panic(err)
	}

	defer db.Close()
}
