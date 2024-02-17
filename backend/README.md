# migration コマンド

テーブル作成

`migrate -path db/migrations -database "mysql://root:root@tcp(localhost:3306)/chat_db" up`

テーブル削除

`migrate -path db/migrations -database "mysql://root:root@tcp(localhost:3306)/chat_db" down  `

テーブル作成

`migrate create -ext sql -dir db/migrations -seq create_users_table `
