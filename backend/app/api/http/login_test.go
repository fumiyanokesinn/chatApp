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

func TestLogin(t *testing.T) {
	config.GetTestEnv()
	test.WithTransaction(t, func(t *testing.T, db *sql.Tx) {

		loginInfo := struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}{
			Email:    "alice@example.com",
			Password: "password",
		}

		requestBody, err := json.Marshal(loginInfo)
		if err != nil {
			t.Fatalf("Error marshaling login info: %v", err)
		}

		router := api.SetRouter()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(requestBody))

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code 200, but got %d", w.Code)
		}
	})
}
