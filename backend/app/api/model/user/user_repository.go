package user

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

// UserRepository はユーザー情報にアクセスするためのインターフェイスです。
type UserRepository interface {
	FindByEmail(email string) (*User, error)
}

// SQLUserRepository はUserRepositoryのSQLに基づく実装です。
type SQLUserRepository struct {
	DB *sql.DB
}

// NewSQLUserRepository は新しいSQLUserRepositoryインスタンスを作成します。
func NewSQLUserRepository(db *sql.DB) UserRepository {
	return &SQLUserRepository{DB: db}
}

// FindByEmail は指定されたメールアドレスに基づいてユーザーを検索します。
func (repo *SQLUserRepository) FindByEmail(email string) (*User, error) {
	var user User
	query := "SELECT id, name, email, password FROM users WHERE email = ?"
	err := repo.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return &user, err
}
