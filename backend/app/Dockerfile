# Goのバージョンを指定
FROM golang:latest

# 作業ディレクトリを設定
WORKDIR /app

# ホストマシンのファイルをコンテナ内の作業ディレクトリにコピー
COPY . .

# 依存関係をダウンロード
RUN go mod download

# アプリケーションをビルド
RUN go build -o main .

# golang-migrateのバージョンを指定
ARG MIGRATE_VERSION=v4.15.1

# golang-migrateのバージョンを指定
ARG MIGRATE_VERSION=v4.15.1

# golang-migrateをダウンロードしてインストール
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz | tar xvz -C /usr/local/bin migrate

# アプリケーションを実行
# CMD ["./main"]
