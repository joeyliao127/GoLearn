//go:build ignore

// homework-03-tag｜搭配 sample-03-tag
// 目標（≤5 分鐘）：幫 Order struct 補上 json tag，讓它序列化成小寫鍵名的 JSON。
// 電商情境：訂單要回給前端 API，鍵名要是 order_id / customer / total（不是預設的大寫欄位名）。
//
// TODO 1：幫 Order 的三個欄位補上 json tag：
//
//	ID → `json:"order_id"`、Customer → `json:"customer"`、Total → `json:"total"`
//	（tag 用反引號，接在欄位型別後面，例如：ID string `json:"order_id"`）
//
// # TODO 2：（已寫好，不用改）main 會用 json.MarshalIndent 把 order 轉成 JSON 並印出
//
// 完成後跑：go run 005-struct與方法/homework-03-tag.go
// 驗收：輸出的 JSON 鍵名為 "order_id"、"customer"、"total"（值為 O-777 / 王小明 / 4980）
package main

import (
	"encoding/json"
	"fmt"
)

// TODO 1：幫下面三個欄位補上 json tag
type Order struct {
	ID       string
	Customer string
	Total    int
}

func main() {
	order := Order{ID: "O-777", Customer: "王小明", Total: 4980}

	b, err := json.MarshalIndent(order, "", "  ")
	if err != nil {
		fmt.Println("序列化失敗：", err)
		return
	}
	fmt.Println(string(b))
}
