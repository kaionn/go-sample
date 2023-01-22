package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// 構造体
type post struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	e.GET("/", hello)
	e.POST("/:id", postSample)

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))
}

// ハンドラーを定義
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!!!")
}

// Handler
func postSample(c echo.Context) error {
	fmt.Println(c)
	p := new(post)
	if err := c.Bind(p); err != nil {
		log.Printf("err %v", err.Error())
		return c.String(http.StatusInternalServerError, "Error!")
	}

	// URLパラメーターはBindで入らない
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("err %v", err.Error())
		return c.String(http.StatusInternalServerError, "Error!")
	}
	p.ID = id
	msg := fmt.Sprintf("id: %v, name %v", p.ID, p.Name)
	return c.String(http.StatusOK, msg)
}
