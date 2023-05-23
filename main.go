package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"time"
	"utils"
)

type URL struct {
	URL string `form:"URL" binding:"required"`
}

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
		Addr:     "127.0.0.1:6379",
		Password: "GTO4mjZQXZkWYgspMWHHgla0Lf5yNew8zlgRyq", // 密码
		DB:       0,                                        // 数据库
		PoolSize: 20,                                       // 连接池大小
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
	//res, err := rdb.Get("a").Result()
	//if err != nil {
	//	panic(err)
	//}

	r := gin.Default()
	// test for ping
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 创建短链接
	r.POST("c", func(c *gin.Context) {
		var b URL
		if err := c.ShouldBind(&b); err != nil {
			c.JSON(200, gin.H{
				"message": "URL required!",
			})
			return
		}
		URL := c.PostForm("URL")
		cacheURL := utils.HashShortURL(URL)
		_, err := rdb.Get(cacheURL).Result()
		if err == redis.Nil {
			//err := rdb.Set(URL, cacheURL, 86400*time.Second).Err()
			err := rdb.Set(cacheURL, URL, 86400*time.Second).Err()
			if err != nil {
				panic(err)
			}
		} else if err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"message": cacheURL,
		})
	})
	//r.GET("redis", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": res,
	//	})
	//})
	// 短地址跳转
	r.GET("s/:HASH", func(c *gin.Context) {
		HASH := c.Param("HASH")
		originURL, err := rdb.Get(HASH).Result()
		if err == redis.Nil {
			c.JSON(404, gin.H{
				"message": "error",
				"data":    originURL,
			})
		} else if err != nil {
			c.JSON(404, gin.H{
				"message": "error1",
				"data":    "Not Found",
			})
		}
		c.Redirect(http.StatusMovedPermanently, originURL)
	})

	r.Run(":8085")
}
