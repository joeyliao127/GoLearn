//go:build ignore

// homework-02-pointer接收者｜搭配 sample-02-pointer接收者
// 目標（≤5 分鐘）：幫 Cart 寫一個「會改到自身」的方法——加購一件商品後、總數 +1。
// 電商情境（購物車已給好）：使用者按「加入購物車」，Cart.Count 要真的加上去。
// 因為方法要改到 Cart 本身，receiver 必須用指標 (*Cart)。
//
// TODO 1：幫 Cart 寫方法 func (c *Cart) Add()，函式內把 c.Count 加 1
// （提示：pointer receiver 才改得到本尊；若寫成 (c Cart) 改的是副本，白改）
// TODO 2：在 main 裡對 cart 呼叫 3 次 Add()，再印出 cart.Count
//
// 完成後跑：go run 006-指標/homework-02-pointer接收者.go
// 驗收：印出「購物車件數：3」（若 receiver 誤用值 (c Cart)，會一直是 0）
package main

import "fmt"

type Cart struct {
	Count int
}

// TODO 1：在這裡幫 Cart 寫 func (c *Cart) Add()，把 c.Count 加 1

func main() {
	cart := &Cart{}

	// 下面這行 _ = ... 只是讓半成品能編譯；寫好 Add 並呼叫它之後就把這行刪掉。
	_ = cart

	// 在這裡開始寫 ↓（TODO 2：呼叫 3 次 cart.Add()，再印出 cart.Count）
	fmt.Println("TODO: 還沒開始，把 TODO 1~2 完成後刪掉這行")
}
