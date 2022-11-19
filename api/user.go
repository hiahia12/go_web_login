package api

import (
	"github.com/gin-gonic/gin"
	"go_web_login/dao"
	"go_web_login/model"
	"net/http"
)

func register(c *gin.Context) {
	u := model.User{}
	u.Name = c.PostForm("username")
	u.Password = c.PostForm("password")
	u.Problem = c.PostForm("question")
	u.Answer = c.PostForm("answer")
	flag := dao.SearchUser(&u)
	if !flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user already exists"})
		return
	}
	dao.Adduser(&u)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "add user successful",
	})

}

func login(c *gin.Context) {
	u := model.User{}
	u.Name = c.PostForm("username")
	u.Password = c.PostForm("password")
	flag := dao.SearchUser(&u)
	if flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user doesn't exists"})
		return
	}
	if dao.Searchpassword(&u) != u.Password {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "wrong password",
		})
		return
	}
	c.SetCookie("gin_demo_cookie", "test", 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "login successful",
	})
}
func searchpassword(c *gin.Context) {
	u := model.User{}
	u.Name = c.PostForm("username")
	if dao.SearchUser(&u) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user doesn't exists"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": dao.Searchquestion(&u),
	})
	u.Answer = c.PostForm("answer")
	if !dao.Check(&u, u.Answer) {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "wrong answer",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": dao.Searchpassword(&u),
	})
}
func change(c *gin.Context) {
	u := model.User{}
	u.Name = c.PostForm("username")
	u.Password = c.PostForm("password")
	newpassword := c.PostForm("newpassword")
	if dao.SearchUser(&u) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user doesn't exists"})
		return
	}
	if dao.Searchpassword(&u) != u.Password {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "wrong password",
		})
		return
	}
	dao.Change(&u, newpassword)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "change successful",
	})
}

func delate(c *gin.Context) {
	u := model.User{}
	u.Name = c.PostForm("username")
	u.Password = c.PostForm("password")
	if dao.SearchUser(&u) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user doesn't exists"})
		return
	}
	if dao.Searchpassword(&u) != u.Password {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "wrong password",
		})
		return
	}
	dao.Delate(&u)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "delete successful",
	})
}
