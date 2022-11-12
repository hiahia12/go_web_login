package main

import (
	"go_web_login/api"
	"go_web_login/dao"
)

func main() {
	dao.InitMap()
	api.InitRouter()
}
