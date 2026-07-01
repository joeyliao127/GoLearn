//go:build ignore

// homework-01-error基礎｜搭配 sample-01-error基礎
// 目標（≤5 分鐘）：寫一個會回傳 error 的 withdraw，餘額不足時把錯誤當「值」交回來。
// 電商情境（輸入已給好）：錢包餘額 balance，這次要扣 amount。
// 規則：amount 大於 balance → 回 error；否則回 (扣完的餘額, nil)。
//
// TODO 1：宣告函式 withdraw(balance, amount int) (int, error)。
// 餘額不足（amount > balance）時，用 fmt.Errorf 回傳 error（訊息帶上 amount 與 balance）；否則回傳 (balance-amount, nil)。
// TODO 2：在 main 用 withdraw(1000, 1500) 呼叫，照慣例 if err != nil 先處理錯誤（印出並 return）；沒錯才印「扣款成功，餘額 X 元」。
//
// 完成後跑：go run 008-錯誤處理/homework-01-error基礎.go
// 驗收：withdraw(1000, 1500) 走錯誤分支，印出餘額不足；改成 withdraw(1000, 300) 應印「餘額 700 元」
package main

import "fmt"

// TODO 1：在這裡宣告 func withdraw(balance, amount int) (int, error)

func main() {
	balance := 1000
	amount := 1500

	// 下面兩行 _ = ... 只是讓半成品能編譯；寫好 withdraw 並呼叫它之後就把這兩行刪掉。
	_ = balance
	_ = amount

	// 在這裡開始寫 ↓（TODO 2：呼叫 withdraw，if err != nil 早 return，否則印成功）
	fmt.Println("TODO: 還沒開始，把 TODO 1~2 完成後刪掉這行")
}
