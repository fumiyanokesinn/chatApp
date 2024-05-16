package model

import (
	"testing"

	"github.com/fumiyanokesinn/chatApp/config"
)

func TestConnectDB(t *testing.T) {
	config.GetTestEnv()

	db := ConnectDB()

	if err := db.Ping(); err != nil {
		t.Errorf("データベースへのPingに失敗しました: %v", err)
	}
}
