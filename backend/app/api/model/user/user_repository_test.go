package user

import (
	"testing"

	"github.com/fumiyanokesinn/chatApp/api/model"
)

func TestFindByEmail(t *testing.T) {
	db := model.ConnectDB()
	userRepo := NewUserRepository(db)

	user, err := userRepo.FindByEmail("alice@example.com")
	if err != nil {
		t.Fatalf("Error getting user: %v", err)
	}

	if user.Name != "Alice" {
		t.Errorf("検索が上手く行ってません。期待される名前: Alice, 検索結果: %s", user.Name)
	}
}

func TestCreateUser(t *testing.T) {
	db := model.ConnectDB()
	userRepo := NewUserRepository(db)

	var userInfo = User{
		Name:     "test",
		Email:    "test@example.com",
		Password: "password",
	}

	user, err := userRepo.CreateUser(userInfo)

	if err != nil {
		t.Fatalf("Error getting user: %v", err)
	}
	if user.ID == 0 {
		t.Errorf("登録が上手く行ってません。IDが0です")
	}
	if user.Name != "test" {
		t.Errorf("登録が上手く行ってません。期待される名前: Alice, 登録結果: %s", user.Name)
	}
	if user.Email != "test@example.com" {
		t.Errorf("登録が上手く行ってません。期待されるメール: test@example.com, 登録結果: %s", user.Email)
	}
	if user.Password != "password" {
		t.Errorf("登録が上手く行ってません。期待されるパスワード: password, 登録結果: %s", user.Password)
	}

}
