package user

import (
	"fmt"

	"github.com/fumiyanokesinn/chatApp/api/model/user"
)

type UserService interface {
	StoreUser(user user.User) (*user.User, error)
}

type userService struct {
	UserRepo user.UserRepository
}

func NewUserService(repo user.UserRepository) *userService {
	return &userService{UserRepo: repo}
}

func (s *userService) StoreUser(userRequest user.User) (*user.User, error) {
	if userRequest.ID == 0 {
		user, _ := s.UserRepo.CreateUser(userRequest)
		return user, nil
	} else {
		fmt.Printf("更新します")
		return nil, nil
	}
}
