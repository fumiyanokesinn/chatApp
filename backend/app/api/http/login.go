package http

import (
	"net/http"

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
	isSuccess, err := auth.Authenticate(loginInfo)

	// 認証に失敗した場合、エラーメッセージを含むレスポンスを返す
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "パスワードが違います" {
			statusCode = http.StatusUnauthorized // 認証エラーの場合は401を返す
		} else {
			statusCode = http.StatusNotFound // ユーザーが見つからない場合は404を返す
		}

		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	if !isSuccess {
		// isSuccessがfalseの場合でもエラーを返す（このケースが発生するかはAuthenticateの実装に依存）
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// 認証が成功した場合、成功メッセージを含むレスポンスを返す
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
