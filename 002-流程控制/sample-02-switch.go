//go:build ignore

// sample-02-switch｜對應「教學.html › 4~6 節」
// 電商情境：依「訂單狀態 OrderStatus」印出對應的處理訊息（出貨中心的看板）。
// 重點看 Go 的 switch：case 自動 break（不像 C/Java 會貫穿）、一個 case 可列多個值、
// 以及「無條件 switch」當漂亮的 if-else 鏈用。
// 跑法：go run 002-流程控制/sample-02-switch.go
package main

import "fmt"

// 沿用 001 學過的 iota 列舉：訂單狀態。
type OrderStatus int

const (
	Pending   OrderStatus = iota // 0 待付款
	Paid                         // 1 已付款
	Shipped                      // 2 已出貨
	Completed                    // 3 已完成
	Cancelled                    // 4 已取消
)

func main() {
	// ── 4. 基本 switch：依值分支，case 自動 break ────────────
	// Go 的 case 執行完「自動跳出」，不用像 C/Java 每個 case 尾巴寫 break。
	status := Paid
	switch status {
	case Pending:
		fmt.Println("待付款：提醒顧客完成付款")
	case Paid:
		fmt.Println("已付款：通知倉庫揀貨") // ← 命中這條，印完就結束，不會漏到下面
	case Shipped:
		fmt.Println("已出貨：推播物流追蹤碼")
	case Completed:
		fmt.Println("已完成：邀請顧客評價")
	default:
		fmt.Println("未知狀態：轉人工客服")
	}

	// ── 5. 一個 case 列多個值 + default ──────────────────────
	// 用逗號把多個值放進同一個 case，命中任一個就執行。
	switch status {
	case Pending, Cancelled:
		fmt.Println("尚未進入出貨流程")
	case Paid, Shipped, Completed:
		fmt.Println("已收款，訂單有效")
	default:
		fmt.Println("狀態異常")
	}

	// ── 6. 無條件 switch（switch true）：當 if-else if 鏈的漂亮寫法 ──
	// switch 後面不接值時，等同 switch true，逐條檢查「哪個 case 為真」。
	// 比一長串 else if 好讀，是 Go 的常見慣例（例如依金額分級）。
	amount := 1800
	switch {
	case amount >= 3000:
		fmt.Println("級距：鑽石單")
	case amount >= 1500:
		fmt.Println("級距：黃金單") // 1800 命中這條就停，不會再比下面
	case amount >= 800:
		fmt.Println("級距：白銀單")
	default:
		fmt.Println("級距：一般單")
	}

	// ── 補充：真的想「貫穿」到下一個 case，得手動寫 fallthrough ──
	// fallthrough 會「無條件」執行下一個 case（不再檢查它的條件），很少用、易踩雷。
	level := 2
	switch level {
	case 2:
		fmt.Println("權限：可退款")
		fallthrough // 故意貫穿：等級 2 也含等級 1 的權限
	case 1:
		fmt.Println("權限：可查訂單")
	}
}
