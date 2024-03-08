package auth

import (
	"fmt"

	"github.com/fumiyanokesinn/chatApp/api/model"
	"github.com/fumiyanokesinn/chatApp/api/model/user"
)

type loginInfo struct {
	Email    string
	Password string
}

func Authenticate(loginInfo loginInfo) (bool, error) {
	db := model.ConnectDB()
	userRepo := user.NewUserRepository(db)

	user, err := userRepo.FindByEmail(loginInfo.Email)

	if err != nil {
		return false, err
	}

	if user.Password != loginInfo.Password {
		return false, fmt.Errorf("パスワードが違います")
	}
	return true, nil
}
