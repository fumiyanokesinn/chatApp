package model

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	once sync.Once // インスタンスの初期化を一度だけ行うための変数
	db   *sql.DB   // データベース接続インスタンス
)

// *sql.DB and *sql.Txの共通インターフェース
type Execer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// ConnectDBはデータベース接続のシングルトンインスタンスを返します。
func ConnectDB() *sql.DB {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", os.Getenv("DATABASE"))
		if err != nil {
			fmt.Println("DBに接続できませんでした")
			panic(err)
		}
	})
	return db
}
