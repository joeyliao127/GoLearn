//go:build ignore

// sample-03-iota列舉｜對應「教學.html › 9~10 節」
// 電商情境：訂單狀態 OrderStatus，用 iota 自動編號，再加 String() 印中文。
// 跑法：go run 001-變數與型別/sample-03-iota列舉.go
package main

import "fmt"

// ── 9. const + iota：Go 沒有 enum，用這招頂替 ─────────────
// 慣用做法：自訂一個整數型別 + iota 自動編號。
// iota = 「這一行在 const 區塊裡的索引（第幾行）」，每個區塊從 0 開始。
// 訣竅：某行把右邊算式整個省略 → 自動沿用上一行寫法，而 iota 每行 +1。
type OrderStatus int

const (
	Pending   OrderStatus = iota // 0 待付款 ← 只在這寫一次 iota
	Paid                         // 1 已付款 ← 右邊省略，自動沿用 "= iota"
	Shipped                      // 2 已出貨
	Completed                    // 3 已完成
	Cancelled                    // 4 已取消
)

// ── 10. 加 String() → 印出來是名字不是數字（這步才「真的像 enum」）──
// fmt 有個約定：型別若有 String() 方法，印它時就自動呼叫它。
func (s OrderStatus) String() string {
	names := []string{"待付款", "已付款", "已出貨", "已完成", "已取消"}
	if s < Pending || s > Cancelled {
		return "未知狀態"
	}
	return names[s]
}

func main() {
	fmt.Printf("Pending 的底層數字=%d，印出來=%v\n", Pending, Pending)

	// 模擬一張訂單的狀態流轉
	status := Pending
	fmt.Printf("下單後：%v\n", status)

	status = Paid
	fmt.Printf("付款後：%v\n", status)

	status = Shipped
	fmt.Printf("出貨後：%v（底層其實是 %d）\n", status, status)

	// %v 會呼叫 String() 印中文；%d 直接印底層整數 —— 兩者對照
	fmt.Printf("用 %%d 看全部：待付款=%d 已付款=%d 已出貨=%d 已完成=%d 已取消=%d\n",
		Pending, Paid, Shipped, Completed, Cancelled)
}
