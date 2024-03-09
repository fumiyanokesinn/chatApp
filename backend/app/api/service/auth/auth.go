package auth

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/fumiyanokesinn/chatApp/api/model"
	"github.com/fumiyanokesinn/chatApp/api/model/user"
	"github.com/gin-gonic/gin"
)

type LoginInfo struct {
	Email    string
	Password string
}

var AuthMessages = map[string]string{
	"NotFoundUser":     "ユーザーが見つかりません。",
	"PasswordMismatch": "パスワードが違います。",
	"ServerError":      "サーバーエラーが発生しました。",
	"Success":          "ログインに成功しました",
}

func Authenticate(loginInfo LoginInfo) error {
	db := model.ConnectDB()
	userRepo := user.NewUserRepository(db)

	user, err := userRepo.FindByEmail(loginInfo.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			// 該当するユーザーが見つからない場合の処理
			return fmt.Errorf(AuthMessages["NotFoundUser"])
		} else {
			// その他のエラーの場合の処理
			return fmt.Errorf(AuthMessages["ServerError"])
		}
	}

	if user.Password != loginInfo.Password {
		return fmt.Errorf(AuthMessages["PasswordMismatch"])
	}

	return nil
}

func HandleAuthError(c *gin.Context, err error) {
	if err != nil {
		switch err.Error() {
		case AuthMessages["NotFoundUser"]:
			c.JSON(http.StatusNotFound, gin.H{"message": AuthMessages["NotFoundUser"]})
		case AuthMessages["PasswordMismatch"]:
			c.JSON(http.StatusUnauthorized, gin.H{"message": AuthMessages["PasswordMismatch"]})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": AuthMessages["ServerError"]})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"message": AuthMessages["Success"]})
	}
}
