package api

import (
	"github.com/gin-gonic/gin"
	"go_web_login/dao"
)

func InitRouter() {
	dao.Init()
	r := gin.Default()
	r.POST("/register", register)
	r.POST("/login", login)
	r.PUT("/change", change)
	r.GET("/search", searchpassword)
	r.DELETE("/delete", delate)
	r.Run(":8088")
}
