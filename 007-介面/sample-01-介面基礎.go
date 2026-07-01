//go:build ignore

// sample-01-介面基礎｜對應「教學.html › 1~4 節」
// 電商情境：訂單成立後要「發通知」給顧客，但通知管道有很多種（Email、簡訊…）。
// 我們用一個 Notifier 介面把「怎麼發」抽象掉，訂單流程只認介面、不認實作。
// 重點：① 介面＝一組方法簽章；② 隱式實作（不用寫 implements，方法對上就算實作）；
// ③ duck typing（會 Send 的都能當 Notifier）；④ 用 %T 看介面底下「真正的型別」。
// 跑法：go run 007-介面/sample-01-介面基礎.go
package main

import "fmt"

// ── 1. 定義介面：一組「方法簽章」的集合 ─────────────────────
// Notifier 只要求一個方法 Send(message)。介面裡「只有簽章、沒有實作」。
// Go 慣例：小介面（常只有一個方法），名字常以 -er 結尾（Reader/Writer/Notifier）。
type Notifier interface {
	Send(message string) error
}

// ── 2. 隱式實作：EmailNotifier「剛好」有 Send 方法，就自動是 Notifier ──
// 注意：沒有任何 `implements Notifier` 的宣告。Go 是「方法對得上就算數」。
type EmailNotifier struct {
	From string
}

// 方法簽章 (message string) error 與介面完全一致 → EmailNotifier 滿足 Notifier。
func (e EmailNotifier) Send(message string) error {
	fmt.Printf("[Email] 從 %s 寄出：%s\n", e.From, message)
	return nil
}

// 另一個實作：SMSNotifier。同樣只是「剛好」有相同簽章的 Send。
type SMSNotifier struct {
	Gateway string
}

func (s SMSNotifier) Send(message string) error {
	fmt.Printf("[SMS] 經由 %s 發送：%s\n", s.Gateway, message)
	return nil
}

func main() {
	// ── 3. duck typing：兩種實作都能塞進同一個 Notifier 變數 ───
	// 「如果它會 Send，那它就能當通知器用」——不管底層是 Email 還是 SMS。
	var n Notifier

	n = EmailNotifier{From: "shop@example.com"} // Email 當 Notifier
	_ = n.Send("您的訂單 #1001 已成立")

	n = SMSNotifier{Gateway: "twilio"} // 換成 SMS，呼叫端一行都不用改
	_ = n.Send("您的訂單 #1001 已出貨")

	// ── 4. 用「介面切片」統一處理多個實作 ────────────────────
	// 把不同實作放進 []Notifier，一個迴圈全部通知——這就是介面的威力：
	// 對「一群長得不一樣、但都會 Send 的東西」寫同一段程式。
	channels := []Notifier{
		EmailNotifier{From: "shop@example.com"},
		SMSNotifier{Gateway: "twilio"},
	}
	fmt.Println("--- 對所有管道廣播 ---")
	for _, c := range channels {
		// %T 印出介面底下「真正的動態型別」，看得到 Email / SMS 的差別。
		fmt.Printf("(動態型別 %T) ", c)
		_ = c.Send("週年慶 5 折起！")
	}
}
