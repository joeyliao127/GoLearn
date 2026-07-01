//go:build ignore

// homework-02-switch｜搭配 sample-02-switch
// 目標（≤5 分鐘）：用 switch 依「付款方式」回傳手續費（元），再印出來。
// 電商情境（輸入已給好）：付款方式 method 是字串。
// 費率規則：credit（信用卡）30；atm（ATM 轉帳）15；cod（貨到付款）50；其它 → 0。
//
// TODO 1：宣告函式 feeOf(method string) int，用 switch method 回傳手續費；credit/atm/cod 各一個 case、其它用 default 回 0
// TODO 2：在 main 裡用 method := "cod" 呼叫 feeOf，並用 fmt.Printf 印「付款方式 X → 手續費 Y 元」
//
// 完成後跑：go run 002-流程控制/homework-02-switch.go
// 驗收：method="cod" 時手續費 50；改成 "credit" 是 30、"atm" 是 15、"line" 之類是 0
package main

import "fmt"

// TODO 1：在這裡宣告 func feeOf(method string) int（用 switch 實作）

func main() {
	method := "cod"

	// 下面這行 _ = ... 只是讓半成品能編譯；寫好 feeOf 並呼叫它之後就把這行刪掉。
	_ = method

	// 在這裡開始寫 ↓（TODO 2：呼叫 feeOf 並印出結果）
	fmt.Println("TODO: 還沒開始，把 TODO 1~2 完成後刪掉這行")
}
