package api

import (
	"github.com/gin-gonic/gin"
	"go_web_login/dao"
)

func InitRouter() {
	dao.Init()
	dao.InitRedis()
	r := gin.Default()
	r.POST("/register", register)
	r.POST("/login", login)
	r.PUT("/change", change)
	r.GET("/search", searchpassword)
	r.DELETE("/delete", delate)
	r.GET("/look", look)
	r.POST("/write", write)
	r.POST("/thumb", thumb)
	r.Run(":8088")

}
