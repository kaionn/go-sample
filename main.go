package main

import (
	"github.com/kaionn/go-sample/src/api/filter_word"
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
	e.GET("/", filter_word.FilterText)

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))
}
