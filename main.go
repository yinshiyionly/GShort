package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

// URL gin 参数验证
type URL struct {
	URL string `form:"URL" binding:"required"`
}

var chars = strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", "")

func encode62(id int64) []int64 {
	var indexAry []int64
	base := int64(len(chars))

	for id > 0 { // i < 0 时,说明已经除尽了,已经到最高位,数字位已经没有了
		remainder := id % base
		indexAry = append(indexAry, remainder)
		id = id / base
	}

	return indexAry
}

// getString62 输出字符串, 长度不一定为6
func getString62(indexAry []int64) string {
	result := ""
	for val := range indexAry {
		result = result + chars[val]
	}

	return reverseString(result[:6])
}

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func main() {
	fmt.Print(encode62(1695629213000000))

	//id := primitive.NewObjectIDFromTimestamp(time.Now()).Hex()
	//fmt.Println(viper.GetString("mongo.database"))
	//var params services.CodeURLItem
	//params.Url = "https://baidu.com"
	//params.Code = "https://baidu.com"

	//fmt.Println(services.InsertToMongo(params))
	// 设置 api 路由
	//router := apiRoute.SetUpApiRouter()
	//// 运行应用程序
	//port := viper.GetString("app.expose")
	//// 设置默认端口号
	//if port == "" {
	//	port = "8085"
	//}
	//err := router.Run("0.0.0.0:" + port)
	//if err != nil {
	//	log.Fatalf("Router run failed : %v", err)
	//}
	//redisClient, _ := services.InitRedis()
	// 注册退出信号处理函数
	//services.SetupGracefulShutdown(redisClient)
	// 当不再需要使用 Redis 客户端时，关闭连接
	//defer redisClient.Close()
	//r := gin.Default()
	//// 创建短链接
	//r.POST("c", func(c *gin.Context) {
	//	var b URL
	//	if err := c.ShouldBind(&b); err != nil {
	//		c.JSON(200, gin.H{
	//			"message": "URL required!",
	//		})
	//		return
	//	}
	//	parseURL, err := url.Parse(b.URL)
	//	if err != nil {
	//		c.JSON(200, gin.H{
	//			"message": "Invalid URL",
	//		})
	//		return
	//	}
	//	if parseURL.Scheme != "http" && parseURL.Scheme != "https" {
	//		c.JSON(200, gin.H{
	//			"message": "Invalid Protocol",
	//		})
	//		return
	//	}
	//	// mongo 中查询
	//	exists, code := services.FindOriginURL(b.URL)
	//	if !exists {
	//		code = services.HashShortURL(b.URL)
	//		// 插入 mongo
	//		params := services.CodeURLItem{Code: code, URL: b.URL}
	//		services.InsertToMongo(params)
	//	}
	//	c.JSON(200, gin.H{
	//		"code": code,
	//	})
	//})
	//// 短地址跳转
	//r.GET("s/:HASH", func(c *gin.Context) {
	//	HASH := c.Param("HASH")
	//	exists, originURL := services.FindOriginURLByCode(HASH)
	//	if exists {
	//		c.Redirect(http.StatusMovedPermanently, originURL)
	//	} else {
	//		c.JSON(404, gin.H{
	//			"message": "error",
	//			"data":    "Not Found",
	//		})
	//	}
	//})
	//r.Run("127.0.0.1:8085")
}

// init 初始化配置文件
func init() {
	// 设置配置文件
	viper.SetConfigFile("./.env.yaml")
	// 设置配置文件类型
	viper.SetConfigType("yaml")
	// 读取配置文件
	err := viper.ReadInConfig()
	// 错误处理
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
}
