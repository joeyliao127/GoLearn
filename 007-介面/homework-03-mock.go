//go:build ignore

// homework-03-mock｜搭配 sample-03-mock
// 目標（≤5 分鐘）：自己寫一個 mock 實作介面，並用它「驗證 service 有沒有真的去發通知」。
// 電商情境：SignupService（註冊服務）成功註冊後，要透過 Notifier 發一封歡迎通知。
// 我們不想真的寄信，改注入一個 MockNotifier，檢查「Send 被呼叫幾次、內容是什麼」。
//
// 已給好：Notifier 介面、SignupService（呼叫 n.Send）。你要補的是那個 mock。
// TODO 1：定義 MockNotifier struct，欄位 Calls int（記錄呼叫次數）、LastMsg string（記錄最後訊息）。
// TODO 2：幫 MockNotifier 寫方法 Send(message string) error，做 Calls++、記下 message、回 nil（這樣就隱式實作了 Notifier）。
// TODO 3：在 main 建 &MockNotifier{}、注入 SignupService、呼叫 Register("小明")，再印「Send 被呼叫 X 次」與「內容：<LastMsg>」。
// 完成後跑：go run 007-介面/homework-03-mock.go
// 驗收：印出「Send 被呼叫 1 次」與「內容：歡迎加入，小明！」
package main

import "fmt"

type Notifier interface {
	Send(message string) error
}

type SignupService struct {
	notifier Notifier
}

func (s *SignupService) Register(name string) {
	// 假裝完成了註冊…然後發歡迎通知（我們要驗證的就是這個呼叫）。
	_ = s.notifier.Send("歡迎加入，" + name + "！")
}

// TODO 1 / 2：在這裡定義 MockNotifier 與它的 Send 方法

func main() {
	// 在這裡開始寫 ↓（TODO 3：建 mock、注入 SignupService、Register("小明")、印出斷言）
	fmt.Println("TODO: 還沒開始，把 TODO 1~3 完成後刪掉這行")
}
