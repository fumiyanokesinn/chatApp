package api

import (
	"github.com/fumiyanokesinn/chatApp/api/http"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	// API動作確認用
	r.GET("/ping", http.Ping)
	// 下にエンドポイントを追加
	return r
}
