package user

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo UserRepository) FindByEmail(email string) (*User, error) {
	var user User

	query := "SELECT id, name, email, password FROM users WHERE email = ?"
	err := repo.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			// 該当するユーザーが見つからない場合の処理
			return nil, fmt.Errorf("ユーザーが見つかりません: %v", email)
		} else {
			// その他のエラーの場合の処理
			return nil, fmt.Errorf("クエリの実行中にエラーが発生しました: %v", err)
		}
	}

	return &user, nil
}
