package http

import (
	"net/http"

	"github.com/fumiyanokesinn/chatApp/api/model/user"
	userService "github.com/fumiyanokesinn/chatApp/api/service/user"
	"github.com/gin-gonic/gin"
)

type CreateAccountHandler struct {
	UserService userService.UserService
}

func NewCreateAccount(userService userService.UserService) *CreateAccountHandler {
	return &CreateAccountHandler{
		UserService: userService,
	}
}

func (h *CreateAccountHandler) CreateAccount(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	userRepo, _ := h.UserService.StoreUser(user)
	c.JSON(http.StatusOK, gin.H{
		"ID":       userRepo.ID,
		"Name":     userRepo.Name,
		"Email":    userRepo.Email,
		"Password": userRepo.Password,
	})
}
