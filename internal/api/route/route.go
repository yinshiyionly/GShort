package apiRoute

import (
	"github.com/gin-gonic/gin"
	apiController "gshort/internal/api/controller"
)

func SetUpApiRouter() *gin.Engine {
	// 创建默认的 Gin 引擎
	router := gin.Default()
	// 定义路由
	apiV1 := router.Group("/api/v1/")
	{
		// 创建短链接
		apiV1.POST("generate", apiController.GenerateTinyUrl)
		// 短链接访问
		apiV1.GET("s", apiController.Visit)
		// 短链接还原
		apiV1.GET("recover", apiController.Recover)
	}

	// 返回路由
	return router
}
