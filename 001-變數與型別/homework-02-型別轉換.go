//go:build ignore

// homework-02-型別轉換｜搭配 sample-02-型別轉換
// 目標（≤5 分鐘）：算出一筆訂單的「含稅總額」與「每件含稅單價」。
// 電商情境（輸入已給好，直接算）：單價 350、數量 4、稅率 0.05（5%）。
//
// TODO 1：算未稅小計 subtotal（int）＝ 單價 × 數量
// TODO 2：算含稅總額 total（float64）＝ subtotal ×（1 + taxRate）；subtotal 是 int，要先 float64(...) 才能混算
// TODO 3：算每件含稅單價（float64）＝ total ÷ quantity；quantity 是 int，想一下要不要也轉 float64？
// TODO 4：用 %.2f 印出「含稅總額」與「每件含稅單價」
//
// 完成後跑：go run 001-變數與型別/homework-02-型別轉換.go
// 驗收：含稅總額 1470.00、每件含稅單價 367.50
package main

import "fmt"

func main() {
	unitPrice := 350
	quantity := 4
	taxRate := 0.05

	// 下面三行 _ = ... 只是讓「還沒動工的半成品」也能編譯；用到某個變數後就把對應那行刪掉。
	_ = unitPrice
	_ = quantity
	_ = taxRate

	// 在這裡開始寫 ↓
	fmt.Println("TODO: 還沒開始，把 TODO 1~4 完成後刪掉這行")
}
