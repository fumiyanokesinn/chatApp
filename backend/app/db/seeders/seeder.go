package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Seederを実行します")
	defer fmt.Println("Seederを終了します")

	db, err := sql.Open("mysql", "root:root@tcp(db:3306)/chat_db")
	if err != nil {
		fmt.Println("DBに接続できませんでした")
		panic(err)
	}
	defer db.Close()

	// Seederを実行
	usersSeeder(db)
}
