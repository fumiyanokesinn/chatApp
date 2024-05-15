package user

import (
	"database/sql"
	"testing"

	"github.com/fumiyanokesinn/chatApp/api/test"
)

func TestFindByEmail(t *testing.T) {
	test.WithTransaction(t, func(t *testing.T, db *sql.Tx) {
		userRepo := NewUserRepository(db)

		// テスト用のユーザーを作成
		_, err := db.Exec("INSERT INTO users (name, email,password) VALUES (?, ?, ?)", "test", "test@example.com", "password")
		if err != nil {
			t.Fatalf("Error creating user: %v", err)
		}

		user, err := userRepo.FindByEmail("test@example.com")
		if err != nil {
			t.Fatalf("Error getting user: %v", err)
		}

		if user.Name != "test" {
			t.Fatalf("検索が上手く行ってません。期待される名前: test, 検索結果: %s", user.Name)
		}
	})
}

func TestCreateUser(t *testing.T) {
	test.WithTransaction(t, func(t *testing.T, db *sql.Tx) {
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
			t.Fatalf("登録が上手く行ってません。IDが0です")
		}
		if user.Name != "test" {
			t.Fatalf("登録が上手く行ってません。期待される名前: Alice, 登録結果: %s", user.Name)
		}
		if user.Email != "test@example.com" {
			t.Fatalf("登録が上手く行ってません。期待されるメール: test@example.com, 登録結果: %s", user.Email)
		}
		if user.Password != "password" {
			t.Fatalf("登録が上手く行ってません。期待されるパスワード: password, 登録結果: %s", user.Password)
		}
	})
}
