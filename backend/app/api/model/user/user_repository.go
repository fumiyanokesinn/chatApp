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

	return &user, err
}
