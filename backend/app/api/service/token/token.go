package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT claims struct
type Claims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

type TokenService interface {
	CreateToken(userID string) (string, error)
}

// tokenService構造体は、TokenServiceインターフェイスを実装します。
type tokenService struct {
}

// NewTokenService関数は、新しいtokenServiceインスタンスを生成します。
func NewTokenService() *tokenService {
	return &tokenService{}
}

// CreateToken generates a JWT token for a given user ID
func (s *tokenService) CreateToken(userID string) (string, error) {
	// Set token claims
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret. Replace "your_secret_key" with your actual secret key.
	signedToken, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
