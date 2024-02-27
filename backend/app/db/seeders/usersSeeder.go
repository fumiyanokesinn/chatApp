package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func usersSeeder(db *sql.DB) {
	users := []struct {
		Name     string
		Email    string
		Password string
	}{
		{"Alice", "alice@example.com", "password"},
		{"Bob", "bob@example.com", "password"},
		{"Charlie", "charlie@example.com", "password"},
	}
	successCount := 0 // 成功した挿入の数をカウント

	// 各ユーザーをデータベースに挿入
	for _, user := range users {
		_, err := db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
		if err != nil {
			fmt.Printf("Failed to insert user: %v\n", err)
		} else {
			successCount++ // 成功カウントを増やす
		}
	}

	if successCount > 0 {
		fmt.Printf("UsersSeeder executed successfully. %d users inserted.\n", successCount)
	} else {
		fmt.Println("Failed to insert any users.")
	}

}
