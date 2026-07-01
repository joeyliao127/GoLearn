//go:build ignore

// homework-01-多回傳值｜搭配 sample-01-多回傳值
// 目標（≤5 分鐘）：寫一個「一次回傳兩個值」的函式，算出「含稅價」與「稅額」。
// 電商情境（輸入已給好）：某商品單價 unitPrice、數量 qty，稅率固定 5%。
//
// TODO 1：宣告函式 taxedPrice(unitPrice, qty int) (total, tax int)（用具名回傳值）；
// 內部算 subtotal = unitPrice*qty、tax = subtotal*5/100、total = subtotal+tax，
// 最後 return total, tax（別用 naked return，明確寫出來）。
// TODO 2：在 main 用多重指派接住 total, tax，再用 fmt.Printf 印「含稅價 X 元（其中稅 Y 元）」。
//
// 完成後跑：go run 003-函式/homework-01-多回傳值.go
// 驗收：unitPrice=1000、qty=2 → 含稅價 2100 元（其中稅 100 元）；改成 qty=1 → 含稅價 1050、稅 50。
package main

import "fmt"

// TODO 1：在這裡宣告 func taxedPrice(unitPrice, qty int) (total, tax int)

func main() {
	unitPrice := 1000
	qty := 2

	// 下面兩行 _ = ... 只是讓半成品能編譯；用到某個變數後就把對應那行刪掉。
	_ = unitPrice
	_ = qty

	// 在這裡開始寫 ↓（TODO 2：呼叫 taxedPrice、接住兩個回傳值、印出來）
	fmt.Println("TODO: 還沒開始，把 TODO 1~2 完成後刪掉這行")
}
