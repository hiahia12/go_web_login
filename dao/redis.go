package dao

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go_web_login/model"
	"log"
	"strconv"
	"time"
)

func InitRedis() {
	model.Rdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		//Password: "123456",
		DB: 0,
	})
	_, err := model.Rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Print(err)
	}
	fmt.Println("redis 链接成功")
}
func Sadd(key string, value int64) error {
	rs := NewRedisSet(context.Background(), key, value, model.Rdb)
	_, err := rs.Conn.SAdd(rs.Context, rs.Object, rs.Id).Result()
	if err != nil {
		return err
	}
	return nil
}
func SetRedisValue(ctx context.Context, key string, value string, expiration time.Duration) error {
	SetKV := model.Rdb.Set(ctx, key, value, expiration)
	return SetKV.Err()
}
func NewRedisSet(context context.Context, Objet string, Id int64, Conn *redis.Client) *model.RedisSet {
	return &model.RedisSet{
		Id:      Id,
		Object:  Objet,
		Conn:    Conn,
		Context: context,
	}
}
func GetRedisValue(ctx context.Context, key string) (string, error) {
	GetKey := model.Rdb.Get(ctx, key)
	if GetKey.Err() != nil {
		return "", GetKey.Err()
	}
	return GetKey.Val(), nil
}
func Hgethallperson(ctx context.Context, key string) (map[string]string, error) {
	GetKey := model.Rdb.HGetAll(ctx, key)
	if GetKey.Err() != nil {
		return nil, GetKey.Err()
	}
	return GetKey.Val(), nil
} //个人信息获取

func Hgethallcheck(ctx context.Context, key1 string, key2 string) (bool, error) {
	GetKey := model.Rdb.HGetAll(ctx, key1)
	fmt.Println("key1=", key1)
	fmt.Println("key2=", key2)
	fmt.Println(GetKey)
	if GetKey.Err() != nil {
		return false, GetKey.Err()
	}
	if GetKey.Val()[key2] == "" {
		return false, nil
	}
	return true, nil
} //检验用户信息是否在redis中
func PasswordaCheck(password string, ctx context.Context, username string) bool {
	user, _ := Hgethallperson(ctx, username)
	if user["password"] != password {
		return false
	}
	return true
} //检验密码是否正确//已排查，暂无问题
func UserDatebaseRedis(user model.User, ctx context.Context) {
	err := Sadd("username", int64(user.Id))
	if err != nil {
		return
	} //先放一个搜索用的？
	userinformation := UserDatebaseMysql(user.Name)
	map1 := map[string]string{
		"name":     userinformation.Name,
		"password": userinformation.Password,
		"id":       string(userinformation.Id),
		"answer":   userinformation.Answer,
		"problem":  userinformation.Problem,
	}
	model.Rdb.HSet(ctx, userinformation.Name, map1)
} //将mysql中人的信息存入redis // 已排查，无问题
func ArticleDatabase(article model.Article, ctx context.Context) {
	Sadd("article", int64(article.Id))
	articleinformation := ArticleDatabaseMysql(article.Id)
	map1 := map[string]string{
		"Id":     string(article.Id),
		"Like":   string(article.Like),
		"Word":   article.Word,
		"Writer": article.Writer,
	}
	model.Rdb.HSet(ctx, string(articleinformation.Id), map1)
}

func Thumb(articleid string, userid string, ctx context.Context) {
	t := "article" + articleid
	int, _ := strconv.Atoi(userid)
	err := Sadd(t, int64(int))
	fmt.Println(err)

}
