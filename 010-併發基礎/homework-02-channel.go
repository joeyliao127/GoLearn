//go:build ignore

// homework-02-channel｜搭配 sample-02-channel
// 目標（≤5 分鐘）：用 channel 把每個 goroutine 算出的小計傳回主流程，加總起來。
// 電商情境（商品清單已給好）：每個商品開一個 goroutine 算 Price×Qty，把小計「送進 channel」；
// 主流程從 channel 收固定筆數（len(items) 筆）並加總，印出購物車總金額。
//
// TODO 1：建一個 channel：results := make(chan int)
// TODO 2：for range items，每個商品用 go func(p Product){ results <- p.Price*p.Qty }(item) 送小計進 channel
// TODO 3：主流程收 len(items) 筆：for i := 0; i < len(items); i++ { total += <-results }，最後印「購物車總金額 X 元」
// 提示：這裡用「收固定筆數」就好，不必 close；close + for range 的寫法見 sample-02。
// 完成後跑：go run 010-併發基礎/homework-02-channel.go
// 驗收：印出「購物車總金額 12660 元」（2490 + 1180 + 8990，收的順序不影響總和）
package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Qty   int
}

func main() {
	items := []Product{
		{Name: "鍵盤", Price: 2490, Qty: 1},
		{Name: "滑鼠", Price: 590, Qty: 2},
		{Name: "螢幕", Price: 8990, Qty: 1},
	}
	total := 0

	// 下面兩行只是讓半成品能編譯；開始用它們之後就把對應那行刪掉。
	_ = items
	_ = total

	// 在這裡開始寫 ↓
	fmt.Println("TODO: 還沒開始，把 TODO 1~3 完成後刪掉這行")
}
