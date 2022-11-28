package model

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var Rdb *redis.Client
var D *gorm.DB

type User struct {
	Id       int    `gorm:"id"`
	Name     string `gorm:"name"`
	Password string `gorm:"password"`
	Problem  string `gorm:"problem"`
	Answer   string `gorm:"answer"`
}
type RedisSet struct {
	Id      int64
	Object  string
	Conn    *redis.Client
	Context context.Context
}
type Article struct {
	Id     int    `gorm:"id"`
	Word   string `gorm:"word"`
	Writer string `gorm:"writer"`
	Like   int    `gorm:"like"`
}
