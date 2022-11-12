package api

import (
	"github.com/gin-gonic/gin"
	"go_web_login/dao"
	"net/http"
)

func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	question := c.PostForm("question")
	answer := c.PostForm("answer")
	if dao.SearchUser(username) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user already exists"})
		return
	}
	dao.Adduser(username, password, question, answer)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "add user successful",
	})
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if !dao.SearchUser(username) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user doesn't exists"})
		return
	}
	if dao.Searchpassword(username) != password {
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
	username := c.PostForm("username")
	if dao.SearchUser(username) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user already exists"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": dao.Searchquestion(username),
	})
	answer := c.PostForm("answer")
	if !dao.Check(username, answer) {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "wrong answer",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": dao.Searchpassword(username),
	})
}
func change(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	newpassword := c.PostForm("newpassword")
	if !dao.SearchUser(username) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user doesn't exists"})
		return
	}
	if dao.Searchpassword(username) != password {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "wrong password",
		})
		return
	}
	dao.Change(username, newpassword)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "change successful",
	})
}

func delate(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if !dao.SearchUser(username) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user doesn't exists"})
		return
	}
	if dao.Searchpassword(username) != password {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "wrong password",
		})
		return
	}
	dao.Delate(username)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "delete successful",
	})
}
