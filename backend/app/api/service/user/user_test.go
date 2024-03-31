package user

import (
	"database/sql"
	"testing"

	"github.com/fumiyanokesinn/chatApp/api/model"
	"github.com/fumiyanokesinn/chatApp/api/model/user"
)

func setup() (*sql.DB, user.UserRepository, *userService) {
	db := model.ConnectDB()
	userRepo := user.NewUserRepository(db)
	userService := NewUserService(userRepo)
	return db, userRepo, userService
}
func TestStoreUser(t *testing.T) {
	_, _, userService := setup()

	var loginInfo = UserInfo{
		Name:     "alice",
		Email:    "alice@example.com",
		Password: "password",
	}

	error := userService.storeUser(loginInfo)

	if error != nil {
		t.Errorf("エラー起きてます")
	}
}
