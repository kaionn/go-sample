package net

import (
	"io"
	"net/http"
)

const PURGO_MALUM_END_POINT string = "https://www.purgomalum.com/service/json"

func FilterByPurgoMalum(word string) string {
	//リクエストを作成 GETの第3引数はnil
	req, _ := http.NewRequest("GET", PURGO_MALUM_END_POINT, nil)
	//requestにheaderをつける
	req.Header.Add("Content-Type", `application/json"`)
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
