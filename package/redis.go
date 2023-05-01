package client

import (
	"fmt"
	"github.com/go-redis/redis"
)

// RedisSingleObj 定义一个 RedisSingleObj 结构体
type RedisSingleObj struct {
	RedisHost     string
	RedisPort     uint16
	RedisAuth     string
	RedisDatabase int
	DB            *redis.Client
}

func (r *RedisSingleObj) InitSingleRedis() (err error) {
	// redis 连接格式拼接
	redisAddr := fmt.Sprintf("%s:%d", r.RedisHost, r.RedisPort)
	// redis 连接对象
	r.DB = redis.NewClient(&redis.Options{
		redisAddr,
		r.RedisAuth,
		r.RedisDatabase,
		300,
		100,
	})
	res, err := r.DB.Ping().Result()
	if err != nil {
		fmt.Printf("Content Failed! err: %v\n", err)
	} else {
		fmt.Printf("Connect Successful! Ping => %v\n", res)
		return nil
	}
	return
}
