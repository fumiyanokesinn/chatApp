package auth

import (
	"testing"
)

func TestAuthenticate(t *testing.T) {
	var loginInfo = LoginInfo{
		Email:    "alice@example.com",
		Password: "password",
	}

	error := Authenticate(loginInfo)

	if error != nil {
		t.Errorf("エラー起きてます")
	}
}

func TestAuthenticateFalseByPassword(t *testing.T) {
	var loginInfo = LoginInfo{
		Email:    "alice@example.com",
		Password: "false",
	}

	err := Authenticate(loginInfo)

	if err.Error() != "パスワードが違います" {
		t.Errorf("期待されるエラーメッセージ: 'パスワードが違います, 実際のエラーメッセージ: '%v'", err.Error())
	}
}
