package http_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fumiyanokesinn/chatApp/api"
	myhttp "github.com/fumiyanokesinn/chatApp/api/http"
	"github.com/gin-gonic/gin"
)

func TestLogin(t *testing.T) {
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
}

func TestHandleLoginError(t *testing.T) {
	// Ginのテスト用のレコーダをセットアップ
	gin.SetMode(gin.TestMode)

	// テストケース
	tests := []struct {
		name            string
		err             error
		expectedCode    int
		expectedMessage string
	}{
		{
			name:            "NotFoundUser error",
			err:             errors.New(myhttp.LoginMessages["NotFoundUser"]),
			expectedCode:    http.StatusNotFound,
			expectedMessage: fmt.Sprintf(`{"message":"%s"}`, myhttp.LoginMessages["NotFoundUser"]),
		},
		{
			name:            "PasswordMismatch error",
			err:             errors.New(myhttp.LoginMessages["PasswordMismatch"]),
			expectedCode:    http.StatusUnauthorized,
			expectedMessage: fmt.Sprintf(`{"message":"%s"}`, myhttp.LoginMessages["PasswordMismatch"]),
		},
		{
			name:            "Unknown error",
			err:             errors.New("Unknown error"),
			expectedCode:    http.StatusInternalServerError,
			expectedMessage: `{"message":"サーバーエラーが発生しました。"}`,
		},
		{
			name:            "No error",
			err:             nil,
			expectedCode:    http.StatusOK,
			expectedMessage: fmt.Sprintf(`{"message":"%s"}`, myhttp.LoginMessages["Success"]),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			myhttp.HandleLoginError(c, tt.err)

			if w.Code != tt.expectedCode {
				t.Errorf("Expected status code %d, but got %d", tt.expectedCode, w.Code)
			}
			if w.Body.String() != tt.expectedMessage {
				t.Errorf("Expected message %s but got %s", tt.expectedMessage, w.Body.String())
			}
		})
	}
}
