package services

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"os/signal"
	"syscall"
)

// 定义一个全部变量
var redisDB *redis.Client

func InitRedis() (*redis.Client, error) {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",                         // 地址
		Password: "GTO4mjZQXZkWYgspMWHHgla0Lf5yNew8zlgRyq", // 密码
		DB:       0,                                        // 数据库
	})
	// 测试连接
	_, err := redisDB.Ping().Result()
	if err != nil {
		return nil, err
	}
	return redisDB, nil
}

// SetupGracefulShutdown 设置优雅退出信号处理
func SetupGracefulShutdown(redisClient *redis.Client) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-stop
		fmt.Println("Shutting down...")

		// 关闭 Redis 连接
		redisClient.Close()

		// 执行其他清理操作...

		os.Exit(0)
	}()
}
