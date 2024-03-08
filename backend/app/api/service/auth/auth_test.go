package auth

import (
	"testing"
)

func TestAuthenticate(t *testing.T) {
	var loginInfo = loginInfo{
		Email:    "alice@example.com",
		Password: "password",
	}

	isSuccess, error := Authenticate(loginInfo)

	if error != nil {
		t.Errorf("エラー起きてます")
	}

	if isSuccess != true {
		t.Errorf("ログインに失敗しました")
	}
}

func TestAuthenticateFalseByPassword(t *testing.T) {
	var loginInfo = loginInfo{
		Email:    "alice@example.com",
		Password: "false",
	}

	isSuccess, err := Authenticate(loginInfo)

	if isSuccess != false {
		t.Errorf("ログインに成功しています")
	}

	if err.Error() != "パスワードが違います" {
		t.Errorf("期待されるエラーメッセージ: 'パスワードが違います, 実際のエラーメッセージ: '%v'", err.Error())
	}
}
