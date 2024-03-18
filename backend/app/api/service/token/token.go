package token

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT claims struct
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type TokenService interface {
	CreateToken(email string) error
	ResponseToken(c *gin.Context, err error)
}

// tokenService構造体は、TokenServiceインターフェイスを実装します。
type tokenService struct {
	Token string
}

// NewTokenService関数は、新しいtokenServiceインスタンスを生成します。
func NewTokenService() *tokenService {
	return &tokenService{}
}

// CreateToken generates a JWT token for a given user ID
func (s *tokenService) CreateToken(email string) error {
	// Set token claims
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return err
	}

	s.Token = signedToken

	return nil
}

func (s *tokenService) ResponseToken(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": s.Token})
}
