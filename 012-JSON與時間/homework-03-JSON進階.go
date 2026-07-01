//go:build ignore

// homework-03-JSON進階｜搭配 sample-03-JSON進階
// 目標（≤5 分鐘）：把一段「結構未知」的 JSON 解進 map[string]any，並用型別斷言取值。
// 電商情境：收到一包第三方金流 webhook，你只想撈出其中幾個欄位來記 log。
//
// TODO 1：把 raw 用 json.Unmarshal 解進 var m map[string]any（記得傳 &m、檢查 err）。
// TODO 2：取出 m["order_id"]（string）與 m["amount"]（注意：JSON 數字會變 float64！）。
// TODO 3：印出「訂單 XXX 金額 YYY 元」；金額用 %.0f 印（因為它是 float64）。
//
// 提示：對照 sample-03-JSON進階 第 11 節；amount 要寫 .(float64)，寫 .(int) 會 panic。
// 完成後跑：go run 012-JSON與時間/homework-03-JSON進階.go
// 驗收：印出「訂單 A-9001 金額 2680 元」
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	raw := []byte(`{"order_id":"A-9001","amount":2680,"paid":true}`)

	// 下面兩行 _ = ... 只是讓半成品能編譯；用到之後就把對應那行刪掉。
	_ = raw
	_ = json.Unmarshal

	// 在這裡開始寫 ↓（TODO 1~3）
	fmt.Println("TODO: 還沒開始，把 TODO 1~3 完成後刪掉這行")
}
