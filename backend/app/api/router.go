package api

import (
	myhttp "github.com/fumiyanokesinn/chatApp/api/http"
	"github.com/fumiyanokesinn/chatApp/api/model"
	"github.com/fumiyanokesinn/chatApp/api/model/user"
	"github.com/fumiyanokesinn/chatApp/api/service/auth"
	"github.com/fumiyanokesinn/chatApp/api/service/token"
	"github.com/fumiyanokesinn/chatApp/config"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	r.Use((config.CORSMiddleware())) // CORSミドルウェアを全体に適用

	db := model.ConnectDB() // データベース接続
	userRepo := user.NewUserRepository(db)
	authService := auth.NewAuthService(userRepo)
	tokenService := token.NewTokenService()
	loginHandler := myhttp.NewLoginHandler(authService, tokenService)

	// API動作確認用
	r.GET("/ping", myhttp.Ping)
	// 下にエンドポイントを追加
	r.POST("/login", loginHandler.Login)
	r.POST("/auth/test", loginHandler.Login)

	return r
}
