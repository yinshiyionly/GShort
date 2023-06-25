package main

import (
	"github.com/gin-gonic/gin"
	"gshort/services"
	"net/http"
	"net/url"
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
		parseURL, err := url.Parse(b.URL)
		if err != nil {
			c.JSON(200, gin.H{
				"message": "Invalid URL",
			})
			return
		}
		if parseURL.Scheme != "http" && parseURL.Scheme != "https" {
			c.JSON(200, gin.H{
				"message": "Invalid Protocol",
			})
			return
		}
		// mongo 中查询
		exists, code := services.FindOriginURL(b.URL)
		if !exists {
			code = services.HashShortURL(b.URL)
		    // 插入 mongo
		    params := services.CodeURLItem{Code: code, URL: b.URL}
		    services.InsertToMongo(params)
		}
		c.JSON(200, gin.H{
			"code": code,
		})
	})
	// 短地址跳转
	r.GET("s/:HASH", func(c *gin.Context) {
		HASH := c.Param("HASH")
		exists, originURL := services.FindOriginURLByCode(HASH)
		if exists {
			c.Redirect(http.StatusMovedPermanently, originURL)
		} else {
			c.JSON(404, gin.H{
				"message": "error",
				"data":    "Not Found",
			})
		}
	})
	r.Run("127.0.0.1:8085")
}
