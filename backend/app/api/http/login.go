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
	err := auth.Authenticate(loginInfo)

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

	// 認証が成功した場合、成功メッセージを含むレスポンスを返す
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

var LoginMessages = map[string]string{
	"NotFoundUser":     "ユーザーが見つかりません。",
	"PasswordMismatch": "パスワードが違います。",
	"Success":          "ログインに成功しました",
}

func HandleLoginError(c *gin.Context, err error) {
	if err != nil {
		switch err.Error() {
		case LoginMessages["NotFoundUser"]:
			c.JSON(http.StatusNotFound, gin.H{"message": LoginMessages["NotFoundUser"]})
		case LoginMessages["PasswordMismatch"]:
			c.JSON(http.StatusUnauthorized, gin.H{"message": LoginMessages["PasswordMismatch"]})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": "サーバーエラーが発生しました。"})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": LoginMessages["Success"]})
	}
}
