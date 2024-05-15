package model

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	once sync.Once // インスタンスの初期化を一度だけ行うための変数
	db   *sql.DB   // データベース接続インスタンス
)

// ConnectDBはデータベース接続のシングルトンインスタンスを返します。
func ConnectDB() *sql.DB {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "root:root@tcp(db:3306)/chat_db")
		if err != nil {
			fmt.Println("DBに接続できませんでした")
			panic(err)
		}
	})
	return db
}

// unit test用のDB接続
func ConnectDBTest() *sql.DB {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "root:root@tcp(db:3306)/chat_db_test")
		if err != nil {
			fmt.Println("DBに接続できませんでした")
			panic(err)
		}
	})
	return db
}
