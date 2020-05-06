package gormte

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
)

func QueryOne(db *gorm.DB) {
	user := User{}
	db.Model(&User{}).Select("name").First(&user)
	Show(user)
}

func QueryPreload(db *gorm.DB) {
	var u = User{}
	// db.Debug().First(&u)
	// db.Debug().Model(&u).Related(&u.Companies)
	// db.Model(&u).Association("Companies").Find(&u.Companies)
	db.Debug().Preload("Companies").Find(&u, 1)
	Show(u)
}

func Show(data interface{}) {
	a, err := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(a), err)
}
