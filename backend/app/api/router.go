package app

import "github.com/gin-gonic/gin"

func SetRouter() *gin.Engine {
	r := gin.Default()
	// API動作確認用
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Health check OK",
		})
	})
	return r
}
