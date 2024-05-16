package user

import (
	"database/sql"
	"testing"

	"github.com/fumiyanokesinn/chatApp/api/model"
	"github.com/fumiyanokesinn/chatApp/api/model/user"
	"github.com/fumiyanokesinn/chatApp/api/test"
	"github.com/fumiyanokesinn/chatApp/config"
)

func setup(db model.Execer) (user.UserRepository, *userService) {
	userRepo := user.NewUserRepository(db)
	userService := NewUserService(userRepo)
	return userRepo, userService
}
func TestStoreUser(t *testing.T) {
	config.GetTestEnv()
	test.WithTransaction(t, func(t *testing.T, db *sql.Tx) {
		_, userService := setup(db)

		var loginInfo = user.User{
			Name:     "alice",
			Email:    "alice@example.com",
			Password: "password",
		}

		_, error := userService.StoreUser(loginInfo)

		if error != nil {
			t.Errorf("エラー起きてます")
		}
	})

}
