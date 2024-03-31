package http

import (
	"net/http"

	"github.com/fumiyanokesinn/chatApp/api/service/user"
	"github.com/gin-gonic/gin"
)

type CreateAccountHandler struct {
}

func NewCreateAccount() *CreateAccountHandler {
	return &CreateAccountHandler{}
}

func (h *CreateAccountHandler) CreateAccount(c *gin.Context) {
	var loginInfo user.UserInfo
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

}
