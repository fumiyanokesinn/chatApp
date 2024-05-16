package http_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fumiyanokesinn/chatApp/api"
	"github.com/fumiyanokesinn/chatApp/api/test"
	"github.com/fumiyanokesinn/chatApp/config"
)

func TestCreateAccount(t *testing.T) {
	config.GetTestEnv()
	test.WithTransaction(t, func(t *testing.T, db *sql.Tx) {
		accountInfo := struct {
			Name     string `json:"name"`
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

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/create_account", bytes.NewBuffer(requestBody))

		// ルーターにリクエストを送信
		router.ServeHTTP(w, req)

		// レスポンスのステータスコードとボディを検証
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code 200, but got %d", w.Code)
		}
	})
}
