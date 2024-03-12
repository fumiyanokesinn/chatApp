package http

import (
	"net/http"

	"github.com/fumiyanokesinn/chatApp/api/service/auth"
	"github.com/fumiyanokesinn/chatApp/api/service/token"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	AuthService  auth.AuthService
	TokenService token.TokenService
}

func NewLoginHandler(authService auth.AuthService, tokenService token.TokenService) *LoginHandler {
	return &LoginHandler{
		AuthService:  authService,
		TokenService: tokenService,
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
