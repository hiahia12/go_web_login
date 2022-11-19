package model

import "github.com/jinzhu/gorm"

type User struct {
	Id       int    `gorm:"id"`
	Name     string `gorm:"name"`
	Password string `gorm:"password"`
	Problem  string `gorm:"problem"`
	Answer   string `gorm:"answer"`
}

var DB *gorm.DB
