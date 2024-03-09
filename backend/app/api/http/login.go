package http

import (
	"net/http"

	"github.com/fumiyanokesinn/chatApp/api/service/auth"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	AuthService auth.AuthService
}

func NewLoginHandler(service auth.AuthService) *LoginHandler {
	return &LoginHandler{
		AuthService: service,
	}
}

func (h *LoginHandler) Login(c *gin.Context) {
	var loginInfo auth.LoginInfo
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.AuthService.Authenticate(loginInfo)
	auth.HandleAuthError(c, err)
}
