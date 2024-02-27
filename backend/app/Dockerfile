# Goのバージョンを指定
FROM golang:latest

# 作業ディレクトリを設定
WORKDIR /app

# ホストマシンのファイルをコンテナ内の作業ディレクトリにコピー
COPY . .

# 依存関係をダウンロード
RUN go mod download

# アプリケーションをビルド２
RUN go build -o main .

# アプリケーションを実行
# CMD ["./main"]
