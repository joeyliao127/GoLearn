//go:build ignore

// homework-01-傳值與指標｜搭配 sample-01-傳值與指標
// 目標（≤5 分鐘）：寫一個「用指標改訂單狀態」的函式，實際改到本尊。
// 電商情境（訂單已給好）：出貨後要把 order.Status 從 "paid" 改成 "shipped"。
// 因為要改到呼叫端的本尊，函式必須收「指標」。
//
// TODO 1：宣告函式 ship(o *Order)，函式內把 o.Status 設成 "shipped"
// （提示：對 *Order 取欄位可直接寫 o.Status，Go 會自動解參考）
// TODO 2：在 main 裡用 & 取地址呼叫 ship(&order)，再印出 order.Status
//
// 完成後跑：go run 006-指標/homework-01-傳值與指標.go
// 驗收：印出「訂單狀態：shipped」（若忘了用指標、寫成 ship(o Order)，狀態會停在 paid）
package main

import "fmt"

type Order struct {
	ID     int
	Status string
}

// TODO 1：在這裡宣告 func ship(o *Order)，把 o.Status 改成 "shipped"

func main() {
	order := Order{ID: 1, Status: "paid"}

	// 下面這行 _ = ... 只是讓半成品能編譯；寫好 ship 並呼叫它之後就把這行刪掉。
	_ = order

	// 在這裡開始寫 ↓（TODO 2：用 ship(&order) 改狀態，再印出 order.Status）
	fmt.Println("TODO: 還沒開始，把 TODO 1~2 完成後刪掉這行")
}
