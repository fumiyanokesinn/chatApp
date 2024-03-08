package user

import (
	"testing"

	"github.com/fumiyanokesinn/chatApp/api/model"
)

func TestNewUserRepository(t *testing.T) {
	db := model.ConnectDB()
	userRepo := NewUserRepository(db)

	if err := userRepo.DB.Ping(); err != nil {
		t.Errorf("データベースへのPingに失敗しました: %v", err)
	}
}

func TestFindByEmail(t *testing.T) {
	db := model.ConnectDB()
	userRepo := NewUserRepository(db)

	user, err := userRepo.FindByEmail("alice@example.com")
	if err != nil {
		t.Fatalf("Error getting user: %v", err) // テストをここで終了させる
	}

	if user.Name != "Alice" {
		t.Errorf("検索が上手く行ってません。期待される名前: Alice, 検索結果: %s", user.Name)
	}
}
