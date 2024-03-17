package token

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func TestCreateToken(t *testing.T) {
	service := NewTokenService()

	email := "alice@example.com"
	err := service.CreateToken(email)

	if err != nil {
		t.Errorf("エラーが発生しました。")
	}

	token, err := jwt.ParseWithClaims(service.Token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
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

	// Emailが期待される値と一致するかを検証
	if claims.Email != email {
		t.Errorf("期待されるEmail(%v)と実際のEmail(%v)が一致しません", email, claims.Email)
	}

}

// 成功時のテスト
func TestResponseTokenSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	s := &tokenService{Token: "testToken"}
	s.ResponseToken(c, nil) // エラーはnil、成功を想定

	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	expectedBody := `{"token":"testToken"}`
	body := w.Body.String()
	if body != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, body)
	}
}

// エラー時のテスト
func TestResponseTokenError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	s := &tokenService{}
	s.ResponseToken(c, errors.New("テスト用エラー")) // someErrorは適当なエラーを想定

	if w.Code != 500 {
		t.Errorf("Expected status code 500, got %d", w.Code)
	}

	expectedBody := `{"error":"Failed to generate token"}`
	body := w.Body.String()
	if body != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, body)
	}
}
