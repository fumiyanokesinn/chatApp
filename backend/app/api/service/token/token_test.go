package token

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
)

func TestCreateToken(t *testing.T) {
	service := NewTokenService()

	userID := "test-user-id"
	tokenString, err := service.CreateToken(userID)

	if err != nil {
		t.Errorf("エラーが発生しました。")
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("your_secret_key"), nil
	})
	if err != nil {
		t.Fatalf("トークンの解析に失敗しました: %v", err)
	}

	// トークンのクレームを検証
	claims, ok := token.Claims.(*Claims)
	if !ok {
		t.Fatal("トークンからクレームを取得できませんでした")
	}

	// トークンの有効性を検証
	if !token.Valid {
		t.Fatal("トークンが無効です")
	}

	// UserIDが期待される値と一致するかを検証
	if claims.UserID != userID {
		t.Errorf("期待されるUserID(%v)と実際のUserID(%v)が一致しません", userID, claims.UserID)
	}

}
