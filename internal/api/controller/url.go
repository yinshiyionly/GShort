package apiController

import (
	"github.com/gin-gonic/gin"
	"gshort/internal/api/model"
	"gshort/services"
	"log"
)

// GenerateTinyUrl 生成短链接
func GenerateTinyUrl(c *gin.Context) {
	var params model.CodeUrlMap
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Fatalf("Gin bind json failed: %v", err)
	}
	services.InsertToMongo(params)
	c.JSON(200, gin.H{
		"message": "success",
	})
	return
}

// Visit  访问
func Visit(c *gin.Context) {

}

// Recover 还原短链接
func Recover(c *gin.Context) {

}
