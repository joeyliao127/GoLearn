//go:build ignore

// homework-01-序列化｜搭配 sample-01-序列化
// 目標（≤5 分鐘）：把一個「商品 struct」轉成 JSON，再把一段 JSON 轉回 struct。
// 電商情境：商品要送給前端（Marshal），也要能接住前端送回來的商品資料（Unmarshal）。
//
// TODO 1：幫 Product 三個欄位補 json tag：Name→"name"、Price→"price"、Stock→"stock,omitempty"（庫存為 0 就不出現在 JSON）。
// TODO 2：用 json.Marshal 把 p 轉成 JSON，印出「商品 JSON: {...}」。（記得檢查 err）
// TODO 3：把 incoming 用 json.Unmarshal 塞進 var got Product（傳 &got），印出「還原後價格: X 元」。
//
// 提示：整套寫法對照 sample-01-序列化；Unmarshal 第二個參數要傳「指標」。
// 完成後跑：go run 012-JSON與時間/homework-01-序列化.go
// 驗收：p 的 JSON 為 {"name":"機械鍵盤","price":2490}（Stock=0 被 omitempty 省略）；還原後價格為 1590 元。
package main

import (
	"encoding/json"
	"fmt"
)

// TODO 1：幫這三個欄位補 json tag
type Product struct {
	Name  string
	Price int
	Stock int
}

func main() {
	p := Product{Name: "機械鍵盤", Price: 2490, Stock: 0}
	incoming := []byte(`{"name":"無線滑鼠","price":1590,"stock":42}`)

	// 下面兩行 _ = ... 只是讓半成品能編譯；用到之後就把對應那行刪掉。
	_ = p
	_ = incoming
	_ = json.Marshal

	// 在這裡開始寫 ↓（TODO 2、TODO 3）
	fmt.Println("TODO: 還沒開始，把 TODO 1~3 完成後刪掉這行")
}
