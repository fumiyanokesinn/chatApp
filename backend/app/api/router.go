package api

import (
	"github.com/fumiyanokesinn/chatApp/api/http"
	"github.com/fumiyanokesinn/chatApp/api/model"
	"github.com/fumiyanokesinn/chatApp/api/model/user"
	"github.com/fumiyanokesinn/chatApp/api/service/auth"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	db := model.ConnectDB() // データベース接続

	userRepo := user.NewSQLUserRepository(db)
	authService := auth.NewAuthService(userRepo)
	loginHandler := http.NewLoginHandler(*authService)

	// API動作確認用
	r.GET("/ping", http.Ping)
	// 下にエンドポイントを追加
	r.POST("/login", loginHandler.Login)
	return r
}
