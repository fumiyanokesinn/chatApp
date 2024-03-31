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
	CreateUser(user User) (*User, error)
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{DB: db}
}

func (repo *userRepository) FindByEmail(email string) (*User, error) {
	var user User
	query := "SELECT id, name, email, password FROM users WHERE email = ?"
	err := repo.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return &user, err
}

func (repo *userRepository) CreateUser(user User) (*User, error) {

	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	result, err := repo.DB.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = int(id)

	return &user, nil
}
