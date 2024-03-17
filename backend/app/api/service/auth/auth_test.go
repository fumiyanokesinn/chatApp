package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fumiyanokesinn/chatApp/api/model"
	"github.com/fumiyanokesinn/chatApp/api/model/user"
	"github.com/gin-gonic/gin"
)

func setup() (*sql.DB, user.UserRepository, *authService) {
	db := model.ConnectDB()
	userRepo := user.NewUserRepository(db)
	authService := NewAuthService(userRepo)
	return db, userRepo, authService
}

func TestAuthenticate(t *testing.T) {
	_, _, authService := setup()

	var loginInfo = LoginInfo{
		Email:    "alice@example.com",
		Password: "password",
	}

	error := authService.Authenticate(loginInfo)

	if error != nil {
		t.Errorf("エラー起きてます")
	}
}

func TestAuthenticateFalseByPassword(t *testing.T) {
	_, _, authService := setup()

	var loginInfo = LoginInfo{
		Email:    "alice@example.com",
		Password: "false",
	}

	err := authService.Authenticate(loginInfo)

	if err.Error() != AuthMessages["PasswordMismatch"] {
		t.Errorf("期待されるエラーメッセージ: '%s, 実際のエラーメッセージ: '%v'", AuthMessages["PasswordMismatch"], err.Error())
	}
}

func TestAuthenticateFalseByEmail(t *testing.T) {
	_, _, authService := setup()

	var loginInfo = LoginInfo{
		Email:    "false@example.com",
		Password: "password",
	}

	err := authService.Authenticate(loginInfo)

	if err.Error() != AuthMessages["NotFoundUser"] {
		t.Errorf("期待されるエラーメッセージ: '%s, 実際のエラーメッセージ: '%v'", AuthMessages["NotFoundUser"], err.Error())
	}
}

func TestHandleAuthError(t *testing.T) {
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
			err:             errors.New(AuthMessages["NotFoundUser"]),
			expectedCode:    http.StatusNotFound,
			expectedMessage: fmt.Sprintf(`{"message":"%s"}`, AuthMessages["NotFoundUser"]),
		},
		{
			name:            "PasswordMismatch error",
			err:             errors.New(AuthMessages["PasswordMismatch"]),
			expectedCode:    http.StatusUnauthorized,
			expectedMessage: fmt.Sprintf(`{"message":"%s"}`, AuthMessages["PasswordMismatch"]),
		},
		{
			name:            "Unknown error",
			err:             errors.New("Unknown error"),
			expectedCode:    http.StatusInternalServerError,
			expectedMessage: `{"message":"サーバーエラーが発生しました。"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			HandleAuthError(c, tt.err)

			if w.Code != tt.expectedCode {
				t.Errorf("Expected status code %d, but got %d", tt.expectedCode, w.Code)
			}
			if w.Body.String() != tt.expectedMessage {
				t.Errorf("Expected message %s but got %s", tt.expectedMessage, w.Body.String())
			}
		})
	}
}
