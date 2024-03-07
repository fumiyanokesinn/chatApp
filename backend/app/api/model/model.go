package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {

	db, err := sql.Open("mysql", "root:root@tcp(db:3306)/chat_db")
	if err != nil {
		fmt.Println("DBに接続できませんでした")
		panic(err)
	}

	return db
}
