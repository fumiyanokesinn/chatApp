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
		t.Fatalf("Error getting user: %v", err)
	}

	if user.Name != "Alice" {
		t.Errorf("検索が上手く行ってません。期待される名前: Alice, 検索結果: %s", user.Name)
	}
}

func TestFindByEmailFalse(t *testing.T) {
	db := model.ConnectDB()
	userRepo := NewUserRepository(db)

	user, err := userRepo.FindByEmail("false@example.com")

	if user != nil {
		t.Errorf("検索が成功しています。取得した名前: %s", user.Name)
	}

	if err.Error() != "ユーザーが見つかりません: false@example.com" {
		t.Errorf("期待されるエラーメッセージ: 'パスワードが違います, 実際のエラーメッセージ: '%v'", err.Error())
	}
}
