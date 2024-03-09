package http

import (
	"net/http"

	"github.com/fumiyanokesinn/chatApp/api/model"
	"github.com/fumiyanokesinn/chatApp/api/model/user"
	"github.com/fumiyanokesinn/chatApp/api/service/auth"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginInfo auth.LoginInfo
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	// ログイン情報の検証処理...
	db := model.ConnectDB()
	userRepo := user.NewSQLUserRepository(db)
	authService := auth.NewAuthService(userRepo)
	err := authService.Authenticate(loginInfo)

	auth.HandleAuthError(c, err)
}
