//go:build ignore

// homework-03-for迴圈｜搭配 sample-03-for迴圈
// 目標（≤5 分鐘）：用 for range 走訪購物車，找出「最貴的一件商品」。
// 電商情境（購物車已給好）：要在結帳頁標示「本單最高價商品」。
//
// TODO 1：宣告 maxPrice（int，先設 0）與 maxName（string，先設 ""）當「目前最貴」的紀錄
// TODO 2：用 for range 走訪 cart，每一項用 if 比較：若 item.Price > maxPrice，就更新 maxPrice 與 maxName
// （提示：range 每輪給 (索引, 元素)，用不到索引就寫 for _, item := range cart）
// TODO 3：迴圈結束後，用 fmt.Printf 印「最貴商品：X（Y 元）」
//
// 完成後跑：go run 002-流程控制/homework-03-for迴圈.go
// 驗收：印出「最貴商品：4K 螢幕（8990 元）」
package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	cart := []Product{
		{Name: "鍵盤", Price: 2490},
		{Name: "4K 螢幕", Price: 8990},
		{Name: "滑鼠", Price: 590},
	}

	// 下面這行 _ = ... 只是讓半成品能編譯；開始用 cart 之後就把這行刪掉。
	_ = cart

	// 在這裡開始寫 ↓
	fmt.Println("TODO: 還沒開始，把 TODO 1~3 完成後刪掉這行")
}
