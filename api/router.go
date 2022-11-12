package api

import "github.com/gin-gonic/gin"

func InitRouter() {

	r := gin.Default()
	r.POST("/register", register)
	r.POST("/login", login)
	r.PUT("/change", change)
	r.DELETE("/delete", delate)
	r.Run(":8088")
}
