package gormte

import (
	"fmt"

	"pkgte/gormte"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var source = "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"

func main() {
	db, err := gorm.Open("mysql", source)
	if err != nil {
		fmt.Println("database init error", err)
	}
	// gormte.Migrate(db)
	gormte.QueryPreload(db)
	defer db.Close()
}
