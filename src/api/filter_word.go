package api

import (
	"net/http"

	"github.com/kaionn/go-sample/src/net/purgo_malum"
	"github.com/labstack/echo"
)

// ハンドラーを定義
func FilterText(c echo.Context) error {
	text := c.QueryParam("text")
	filtered_text := purgo_malum.FilterByPurgoMalum(text)
	return c.String(http.StatusOK, filtered_text)
}
