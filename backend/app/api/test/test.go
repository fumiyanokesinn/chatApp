package test

import (
	"database/sql"
	"testing"

	"github.com/fumiyanokesinn/chatApp/api/model"
	_ "github.com/go-sql-driver/mysql"
)

// setupDBはデータベース接続を設定する関数
func setupDB() *sql.DB {
	return model.ConnectDB()
}

// withTransactionはテスト関数をトランザクション内で実行し、終了後にロールバックする
func WithTransaction(t *testing.T, testFunc func(*testing.T, *sql.Tx)) {
	db := setupDB()

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	// テスト関数を実行
	testFunc(t, tx)
}
