package api

import (
	"net/http"

	myhttp "github.com/fumiyanokesinn/chatApp/api/http"
	"github.com/fumiyanokesinn/chatApp/api/middleware"
	"github.com/fumiyanokesinn/chatApp/api/model"
	"github.com/fumiyanokesinn/chatApp/api/model/user"
	"github.com/fumiyanokesinn/chatApp/api/service/auth"
	"github.com/fumiyanokesinn/chatApp/api/service/token"
	userService "github.com/fumiyanokesinn/chatApp/api/service/user"
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
	userService := userService.NewUserService(userRepo)
	createAccountHandler := myhttp.NewCreateAccount(userService)

	// API動作確認用
	r.GET("/ping", myhttp.Ping)
	// 下にエンドポイントを追加
	r.POST("/login", loginHandler.Login)
	r.POST("/create_account", createAccountHandler.CreateAccount)

	// JWT認証を適用するグループ
	authRequired := r.Group("/api")
	authRequired.Use(middleware.AuthMiddleware())
	{
		authRequired.GET("/secure", func(c *gin.Context) {
			// JWT認証が必要なエンドポイント
			c.JSON(http.StatusOK, gin.H{"message": "Secure content"})
		})
	}

	return r
}
