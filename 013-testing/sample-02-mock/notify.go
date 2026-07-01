// Package notify｜對應「教學.html › 群組 B（3 節：用介面 mock 依賴）」
// 電商情境：新客註冊完要寄一封歡迎通知。但「怎麼送」（Email/SMS/站內信）不該寫死在邏輯裡，
// 所以把它抽成 Notifier 介面——這樣測試時就能塞一個「假的 Notifier」來觀察有沒有被正確呼叫。
// 呼應主題 007：介面在「使用方」定義、小介面（常單方法）、隱式實作。
package notify

// Notifier 是「送通知」的抽象。任何有 Send(string) error 的型別都自動算實作它（duck typing）。
type Notifier interface {
	Send(msg string) error
}

// Welcome 是被測對象：組出歡迎詞，交給注入進來的 Notifier 送出。
// 注意它依賴的是「介面」而不是某個具體寄信器，所以正式跑用真的、測試時用假的。
func Welcome(n Notifier, name string) error {
	return n.Send("歡迎 " + name)
}
