package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go_web_login/model"
)

func Adduser(u *model.User) {
	model.DB.Save(u)
}
func Addarticle(u *model.Article) {
	model.D.Save(u)
}
func SearchUser(u *model.User) bool {
	var t = model.User{}
	model.DB.Where("name=?", u.Name).Debug().First(&t)
	if t.Name == u.Name {
		return false
	}
	return true
}

func Searchpassword(u *model.User) string {
	var t = model.User{}
	model.DB.Where("name=?", u.Name).Debug().First(&t)
	return t.Password
}
func Change(u *model.User, password string) {
	u.Password = password
	model.DB.Model(u).Where("name=?", u.Name).Update("password", password)
}

func Check(u *model.User, answer string) bool {
	var t = model.User{}
	model.DB.Where("name=?", u.Name).Debug().First(&t)
	if t.Answer == answer {
		return true
	}
	return false
}

func Searchquestion(u *model.User) string {
	var t = model.User{}
	model.DB.Where("name=?", u.Name).Debug().First(&t)
	return t.Problem
}

func Delate(u *model.User) {
	model.DB.Where("name=?", u.Name).Debug().Delete(u)
}
func Init() {
	// 链接数据库
	dsn := "root:liao20031103@tcp(127.0.0.1:3306)/learning?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("[WARNING] an error occurred:" + err.Error())
	}
	db.SingularTable(true)
	db.AutoMigrate(&model.User{})
	model.DB = db
	db.AutoMigrate(&model.Article{})
	model.D = db
}
func UserDatebaseMysql(username string) model.User {
	var t = model.User{}
	model.DB.Where("name=?", username).Debug().First(&t)
	return t
} //已排查，无问题
func ArticleDatabaseMysql(articleid int) model.Article {
	var t = model.Article{}
	model.D.Where("id=?", articleid).Debug().First(&t)
	return t
} //已排查，暂无问题？
