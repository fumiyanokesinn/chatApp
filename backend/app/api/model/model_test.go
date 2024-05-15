package model

import (
	"testing"
)

func TestConnectDB(t *testing.T) {
	db := ConnectDB()

	if err := db.Ping(); err != nil {
		t.Errorf("データベースへのPingに失敗しました: %v", err)
	}
}

func TestConnectDBTest(t *testing.T) {
	db := ConnectDBTest()

	if err := db.Ping(); err != nil {
		t.Errorf("データベースへのPingに失敗しました: %v", err)
	}
}
