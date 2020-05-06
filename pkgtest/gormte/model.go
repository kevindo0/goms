package gormte

import (
    "time"
    "github.com/jinzhu/gorm"
)

type User struct {
        ID        int        `gorm:"TYPE:int(11);NOT NULL;PRIMARY_KEY;INDEX"`
        Name      string     `gorm:"TYPE: VARCHAR(255); DEFAULT:'';INDEX"`
        Companies []Company  `gorm:"FOREIGNKEY:UserId;ASSOCIATION_FOREIGNKEY:ID"`
        CreatedAt time.Time  `gorm:"TYPE:DATETIME"`
        UpdatedAt time.Time  `gorm:"TYPE:DATETIME"`
        DeletedAt *time.Time `gorm:"TYPE:DATETIME;DEFAULT:NULL"`
}

type Company struct {
        gorm.Model
        Industry int    `gorm:"TYPE:INT(11);DEFAULT:0"`
        Name     string `gorm:"TYPE:VARCHAR(255);DEFAULT:'';INDEX"`
        Job      string `gorm:"TYPE:VARCHAR(255);DEFAULT:''"`
        UserId   int    `gorm:"TYPE:int(11);NOT NULL;INDEX"`
}
