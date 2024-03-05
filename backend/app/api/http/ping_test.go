package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fumiyanokesinn/chatApp/api"
)

func TestPing(t *testing.T) {
	router := api.SetRouter()
	// テスト用のHTTPリクエストを作成し、`/ping`エンドポイントに対するGETリクエストをシミュレート
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	// ルーターにリクエストを送信
	router.ServeHTTP(w, req)

	// レスポンスのステータスコードとボディを検証
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, but got %d", w.Code)
	}
}
