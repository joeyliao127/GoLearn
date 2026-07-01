//go:build ignore

// homework-01-goroutine｜搭配 sample-01-goroutine
// 目標（≤5 分鐘）：用 goroutine 並行算「每個商品的小計」，並用 WaitGroup 等全部算完。
// 電商情境（商品清單已給好）：要並行處理購物車，每個商品各開一個 goroutine 算 Price×Qty，
// 把小計「印出來」；主流程等全部算完後印一行「全部小計計算完成」。
// （注意：這題只要各自印出小計就好，不用加總 —— 加總會跨 goroutine 共享變數，那是 homework-03 的 Mutex 主題。）
//
// TODO 1：宣告 var wg sync.WaitGroup
// TODO 2：for range items，每個商品先 wg.Add(1)，再 go func(p Product){ defer wg.Done(); ... }(item) 算 p.Price*p.Qty 並 Printf 印「商品 X 小計 Y 元」
// TODO 3：迴圈後 wg.Wait()，再印一行「全部小計計算完成」
//
// 完成後跑：go run 010-併發基礎/homework-01-goroutine.go
// 驗收：會印出 3 行小計（鍵盤 2490、滑鼠 1180、螢幕 8990；順序可能不同），最後一行「全部小計計算完成」
package main

import (
	"fmt"
	"sync"
)

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

	// 下面兩行只是讓半成品能編譯；開始用它們之後就把對應那行刪掉。
	_ = items
	var _ sync.WaitGroup

	// 在這裡開始寫 ↓
	fmt.Println("TODO: 還沒開始，把 TODO 1~3 完成後刪掉這行")
}
