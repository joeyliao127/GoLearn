//go:build ignore

// homework-04-select｜搭配 sample-04-select
// 目標（≤5 分鐘）：用 select 等「付款結果」channel，並加一個 time.After 逾時分支。
// 電商情境：送出付款請求後，等金流回傳結果；但最多只等 200ms，逾時就顯示「稍後查詢」。
// paymentResult channel 已幫你在 50ms 後送回 "付款成功"，所以正常情況會走成功分支。
//
// TODO 1：寫一個 select，同時等兩個 case：
// case msg := <-paymentResult: 印「付款結果：msg」
// case <-time.After(200 * time.Millisecond): 印「付款逾時，稍後查詢」
// TODO 2：（觀察用）跑一次；再把 paymentResult 的 50ms 改大成 300ms，會改走逾時分支 —— 體會 select 選「先就緒」的那個。
//
// 完成後跑：go run 010-併發基礎/homework-04-select.go
// 驗收：印出「付款結果：付款成功」（因為 50ms < 200ms，成功的 channel 先就緒）
package main

import (
	"fmt"
	"time"
)

func main() {
	// 模擬金流：50ms 後把結果送進這個 channel（緩衝 1，送完 goroutine 就能結束）。
	paymentResult := make(chan string, 1)
	go func() {
		time.Sleep(50 * time.Millisecond) // 想驗證逾時分支就把這裡改成 300ms
		paymentResult <- "付款成功"
	}()

	// 下面這行只是讓半成品能編譯；寫好 select 之後就把這行刪掉。
	_ = paymentResult

	// 在這裡開始寫 ↓（TODO 1：用 select 等 paymentResult 或 time.After 逾時）
	fmt.Println("TODO: 還沒開始，把 TODO 1 完成後刪掉這行")
}
