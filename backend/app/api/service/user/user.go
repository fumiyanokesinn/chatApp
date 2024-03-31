package user

import (
	"fmt"

	"github.com/fumiyanokesinn/chatApp/api/model/user"
)

type UserInfo struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type UserService interface {
	storeUser(userInfo UserInfo) error
}

type userService struct {
	UserRepo user.UserRepository
}

func NewUserService(repo user.UserRepository) *userService {
	return &userService{UserRepo: repo}
}

func (s *userService) storeUser(userInfo UserInfo) error {
	if userInfo.ID != 0 {
		fmt.Printf("更新します")
	} else {
		fmt.Printf("登録します")
	}
	return nil
}
