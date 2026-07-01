//go:build ignore

// homework-02-方法｜搭配 sample-02-方法
// 目標（≤5 分鐘）：幫 Product 寫一個 value receiver 方法 Label()，回傳顯示用字串。
// 電商情境：商品列表每一列要顯示「商品名（NT$價格）」，把這段邏輯綁在 Product 上。
//
// TODO 1：宣告方法 func (p Product) Label() string，
//
//	回傳格式為「Name（NT$Price）」，例如「機械鍵盤（NT$2490）」
//	（提示：用 fmt.Sprintf("%s（NT$%d）", p.Name, p.Price)）
//
// # TODO 2：在 main 裡對 keyboard 呼叫 Label()，用 fmt.Println 印出結果
//
// 完成後跑：go run 005-struct與方法/homework-02-方法.go
// 驗收：印出「機械鍵盤（NT$2490）」
package main

import "fmt"

type Product struct {
	ID    string
	Name  string
	Price int
}

// TODO 1：在這裡宣告 func (p Product) Label() string

func main() {
	keyboard := Product{ID: "P001", Name: "機械鍵盤", Price: 2490}

	// 下面這行 _ = ... 只是讓半成品能編譯；寫好 Label 並呼叫它之後就把這行刪掉。
	_ = keyboard

	// 在這裡開始寫 ↓（TODO 2：呼叫 keyboard.Label() 並印出來）
	fmt.Println("TODO: 還沒開始，把 TODO 1~2 完成後刪掉這行")
}
