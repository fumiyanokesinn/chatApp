package api

import (
	"net/http"

	myhttp "github.com/fumiyanokesinn/chatApp/api/http"
	"github.com/fumiyanokesinn/chatApp/api/model"
	"github.com/fumiyanokesinn/chatApp/api/model/user"
	"github.com/fumiyanokesinn/chatApp/api/service/auth"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	r.Use(CORSMiddleware()) // CORSミドルウェアを全体に適用

	db := model.ConnectDB() // データベース接続
	userRepo := user.NewUserRepository(db)
	authService := auth.NewAuthService(userRepo)
	loginHandler := myhttp.NewLoginHandler(authService)

	// API動作確認用
	r.GET("/ping", myhttp.Ping)
	// 下にエンドポイントを追加
	r.POST("/login", loginHandler.Login)
	return r
}

// CORSミドルウェアを定義
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
