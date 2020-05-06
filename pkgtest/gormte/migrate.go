package gormte

import (
    "github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
    db.AutoMigrate(&User{}, &Company{})
}
