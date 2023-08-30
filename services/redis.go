package services

import (
	"fmt"

	redis "github.com/go-redis/redis"
)

var redisDb *redis.Client

func initRedis() (err error) {
	redisDb = redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

	_, err = redisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func GetRedisDb() (db *redis.Client) {
	return redisDb
}

func init() {
	fmt.Println("redis 初始化")
	err := initRedis()
	if err != nil {
		panic(err)
	}
}
