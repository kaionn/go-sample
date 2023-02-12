package main

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	e.GET("/", FilterText)

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))
}

const PURGO_MALUM_END_POINT string = "https://www.purgomalum.com/service/json"

func FilterByPurgoMalum(word string) string {
	//リクエストを作成 GETの第3引数はnil
	req, _ := http.NewRequest("GET", PURGO_MALUM_END_POINT, nil)
	//requestにheaderをつける
	// req.Header.Add("Content-Type", `application/json"`)
	//URLのクエリを確認
	q := req.URL.Query()
	//クエリを追加
	q.Add("text", word)
	//encodeしてからURLに戻す
	req.URL.RawQuery = q.Encode()
	//クライアントを作成
	var client *http.Client = &http.Client{}
	//結果
	resp, _ := client.Do(req)
	//読み込み
	body, _ := io.ReadAll(resp.Body)
	//出力
	return string(body)
}

// ハンドラーを定義
func FilterText(c echo.Context) error {
	text := c.QueryParam("text")
	filtered_text := FilterByPurgoMalum(text)
	return c.String(http.StatusOK, filtered_text)
}
