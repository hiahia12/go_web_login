package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_web_login/dao"
	"go_web_login/model"
	"net/http"
	"strconv"
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
func thumb(c *gin.Context) {
	user := "user"
	articlename := "article"
	u := model.User{}
	u.Name = c.PostForm("username")
	u.Password = c.PostForm("password")
	articleid := c.PostForm("articleid")
	flag1, _ := dao.Hgethallcheck(context.Background(), user, u.Name)
	if !flag1 {
		if dao.SearchUser(&u) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": "user doesn't exists"})
			return
		} //检查用户是否存在redis中
		if dao.Searchpassword(&u) != u.Password {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": "wrong password",
			})
			return
		}
		userimformation := dao.UserDatebaseMysql(u.Name)
		dao.UserDatebaseRedis(userimformation, context.Background()) //先查询redis，redis中无则查询mysql，同时将mysql中值放入redis
	}
	flag := dao.PasswordaCheck(u.Password, context.Background(), u.Name)
	if !flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "wrong password",
		})
		return
	}
	flag2, err := dao.Hgethallcheck(context.Background(), articlename, articleid)
	fmt.Println(err)
	if !flag2 {
		if dao.SearchUser(&u) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": "article doesn't exists"})
			return
		} //检查文章是否存在于redis中
		intarticleid, err := strconv.Atoi(articleid)
		if err != nil {
			print(err)
			return
		}
		articleinformation := dao.ArticleDatabaseMysql(intarticleid)
		dao.ArticleDatabase(articleinformation, context.Background())
	}
	person, _ := dao.Hgethallperson(context.Background(), u.Name)
	dao.Thumb(articleid, person["id"], context.Background())
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  200,
		"message": "点赞成功"})
	return
}
func write(c *gin.Context) {
	u := model.User{}
	u.Name = c.PostForm("username")
	u.Password = c.PostForm("password")
	a := model.Article{}
	a.Word = c.PostForm("word")
	a.Writer = u.Name
	user := "user"
	flag1, _ := dao.Hgethallcheck(context.Background(), user, u.Name)
	if !flag1 {
		if dao.SearchUser(&u) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": "user doesn't exists"})
			return
		} //检查用户是否存在redis中
		userimformation := dao.UserDatebaseMysql(u.Name)
		dao.UserDatebaseRedis(userimformation, context.Background())
	}
	flag := dao.PasswordaCheck(u.Password, context.Background(), u.Name)
	if !flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "wrong password",
		})
		return
	}
	dao.Addarticle(&a)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "write successfully",
	})
} //写留言
func look(c *gin.Context) {
	var a = model.Article{}
	id := c.PostForm("id")
	model.D.Where("id=?", id).Debug().Find(&a)
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  200,
		"message": a,
	})
} //看留言
