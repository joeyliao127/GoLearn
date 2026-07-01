//go:build ignore

// homework-01-介面基礎｜搭配 sample-01-介面基礎
// 目標（≤5 分鐘）：自己定義一個 PaymentMethod 介面，寫「兩個」實作，體會隱式實作與 duck typing。
// 電商情境：結帳時要付款，付款方式有很多種（信用卡、貨到付款…），我們用介面統一「怎麼付」。
//
// TODO 1：定義介面 PaymentMethod，含一個方法 Pay(amount int) string（回傳一句付款說明）
// TODO 2：定義 CreditCard struct（欄位 Last4 string），寫 Pay 方法回傳 "信用卡(末四碼 XXXX) 付款 N 元"
// TODO 3：定義 CashOnDelivery struct（空的即可），寫 Pay 方法回傳 "貨到付款 N 元"
// TODO 4：在 main 宣告 var p PaymentMethod，分別指派兩種實作各呼叫一次 Pay，印出回傳字串
// （提示：不用寫 implements；只要方法簽章 Pay(int) string 對上，就自動是 PaymentMethod）
//
// 完成後跑：go run 007-介面/homework-01-介面基礎.go
// 驗收：印出「信用卡(末四碼 4242) 付款 1500 元」與「貨到付款 1500 元」兩行
package main

import "fmt"

// TODO 1：在這裡定義 PaymentMethod 介面

// TODO 2 / 3：在這裡定義 CreditCard、CashOnDelivery 及各自的 Pay 方法

func main() {
	amount := 1500

	// 下面這行 _ = ... 只是讓半成品能編譯；開始用 amount 之後就把這行刪掉。
	_ = amount

	// 在這裡開始寫 ↓（TODO 4：用 var p PaymentMethod 指派兩種實作並呼叫 Pay）
	fmt.Println("TODO: 還沒開始，把 TODO 1~4 完成後刪掉這行")
}
