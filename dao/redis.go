package dao

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go_web_login/model"
	"time"
)

func InitRedis() {
	model.Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	})
}
func Sadd(key string, value int64) {
	rs := NewRedisSet(context.Background(), key, value, model.Rdb)
	_, err := rs.Conn.SAdd(rs.Context, rs.Object, rs.Id).Result()
	if err != nil {
		fmt.Println(err)
	}
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
	if GetKey.Err() != nil {
		return false, GetKey.Err()
	}
	if GetKey.Val()[key2] == "" {
		return false, nil
	}
	return true, nil
} //检验用户信息是否存在redis
func UserDatebaseRedis(user model.User, ctx context.Context) {
	Sadd("username", int64(user.Id)) //先放一个搜索用的？
	userimformation := UserDatebaseMysql(user.Name)
	map1 := map[string]string{
		"name":     userimformation.Name,
		"password": userimformation.Password,
		"id":       string(userimformation.Id),
		"answer":   userimformation.Answer,
		"problem":  userimformation.Problem,
	}

	model.Rdb.HSet(ctx, userimformation.Name, map1)

} //将mysql中人的信息存入redis
