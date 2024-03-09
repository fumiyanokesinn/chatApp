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

type UserRepository interface {
	FindByEmail(email string) (*User, error)
}

type SQLUserRepository struct {
	DB *sql.DB
}

func NewSQLUserRepository(db *sql.DB) *SQLUserRepository {
	return &SQLUserRepository{DB: db}
}

func (repo *SQLUserRepository) FindByEmail(email string) (*User, error) {
	var user User
	query := "SELECT id, name, email, password FROM users WHERE email = ?"
	err := repo.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return &user, err
}
