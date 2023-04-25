package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func main() {
	//r := gin.Default()
	//// test for ping
	//r.GET("ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.POST("short", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "success",
	//		"data":    "https://t.local",
	//	})
	//})
	//conn := RedisSingleObj{
	//	"101.42.137.30",
	//	6379,
	//	"",
	//}
	//err := conn.InitSingleRedis()
	//if err != nil {
	//	panic(err)
	//}
	//defer conn.DB.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "101.42.137.30:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	// test connect
	//res, err := rdb.Ping().Result()
	//if err != nil {
	//	fmt.Printf("Content Failed! err: %v\n", err)
	//} else {
	//	fmt.Printf("Connect Successful! Ping => %v\n", res)
	//}
	//err := rdb.Set("a", "111", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	res, err := rdb.Get("a").Result()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	// test for ping
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("redis", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": res,
		})
	})
	r.POST("short", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
			"data":    "https://t.local",
		})
	})
	r.Run()
}
