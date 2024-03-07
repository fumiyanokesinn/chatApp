# Goのバージョンを指定
FROM golang:latest

# 作業ディレクトリを設定
WORKDIR /app

# ホストマシンのファイルをコンテナ内の作業ディレクトリにコピー
COPY . .

RUN go mod tidy

# 依存関係をダウンロード
RUN go mod download

# アプリケーションをビルド
RUN go build -o main .

# golang-migrateのバージョンを指定
ARG MIGRATE_VERSION=v4.15.1

# golang-migrateをダウンロードしてインストール
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz | tar xvz -C /usr/local/bin migrate

# go.modでは管理しない、開発に必要なツールをインストール
RUN go install github.com/ramya-rao-a/go-outline@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest


# アプリケーションを実行
# CMD ["./main"]
