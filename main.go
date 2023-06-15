package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gshort/services"
	"net/http"
	"time"
)

type URL struct {
	URL string `form:"URL" binding:"required"`
}

func main() {
	redisClient, _ := services.InitRedis()
	// 注册退出信号处理函数
	services.SetupGracefulShutdown(redisClient)
	// 当不再需要使用 Redis 客户端时，关闭连接
	defer redisClient.Close()
	fmt.Println(services.FindOriginURL("baidu.com"))
	panic(11)

	r := gin.Default()

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
		// mongo 中查询
		cacheURL := services.HashShortURL(URL)
		_, err := redisClient.Get(cacheURL).Result()
		if err == redis.Nil {
			//err := redisClient.Set(URL, cacheURL, 86400*time.Second).Err()
			err := redisClient.Set(cacheURL, URL, 86400*time.Second).Err()
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
		originURL, err := redisClient.Get(HASH).Result()
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
