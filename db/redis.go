package db

import (
	"github.com/go-redis/redis"
)

var dbRedis *redis.Client

func InitRedis() {
	dbRedis = redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})

	_, err := dbRedis.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func GetRedis() *redis.Client {
	return dbRedis
}
