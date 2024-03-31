package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fumiyanokesinn/chatApp/api"
)

func TestCreateAccount(t *testing.T) {
	accountInfo := struct {
		Name     string `json:"string"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Name:     "alice",
		Email:    "alice@example.com",
		Password: "password",
	}
	requestBody, err := json.Marshal(accountInfo)
	if err != nil {
		t.Fatalf("Error marshaling account info: %v", err)
	}
	router := api.SetRouter()
	// テスト用のHTTPリクエストを作成し、`/ping`エンドポイントに対するGETリクエストをシミュレート
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/create_account", bytes.NewBuffer(requestBody))

	// ルーターにリクエストを送信
	router.ServeHTTP(w, req)

	// レスポンスのステータスコードとボディを検証
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, but got %d", w.Code)
	}
}
