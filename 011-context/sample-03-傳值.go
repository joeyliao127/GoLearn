//go:build ignore

// sample-03-傳值｜對應「教學.html › 8~10 節」
// 電商情境：每個進來的請求都帶一個「requestID」（給 log 追蹤用）、有時還有登入的會員 ID。
// 這種「跟著這個請求走、每層都可能想讀」的值，不方便一路加進每個函式參數，
// 就放進 context 裡帶著走——這正是 context.WithValue 的用途。
// 重點：
//  1. WithValue 派生出「帶一個 key/value 的 ctx」。
//  2. key 要用「自訂的未匯出型別」，不要用字串，避免跟別的套件撞 key。
//  3. 讀值用 ctx.Value(key)，回傳 any，要「型別斷言」回原本型別，並處理「拿不到」。
//  4. 慣例：只放 request-scoped 的東西（requestID、使用者身分…），別拿它傳「函式該用參數傳的正經輸入」（如訂單金額、DB 連線）。
//
// 跑法：go run 011-context/sample-03-傳值.go
package main

import (
	"context"
	"fmt"
)

// ── 8. key 用「自訂未匯出型別」，不要用字串 ────────────────────
// 為什麼？context 是全程式共用的，如果大家都用字串 "requestID" 當 key，
// 不同套件很容易「不小心用了同一個字串」而互相覆蓋。
// 用一個「只有本套件看得到的私有型別」當 key，就從型別層面保證不會撞。
type ctxKey int // 未匯出（小寫開頭），外部套件拿不到，也就無法覆蓋我們的值

const (
	keyRequestID ctxKey = iota // 用 iota 給每個 key 一個獨一無二的常數
	keyUserID
)

func main() {
	// 根 ctx（同前：收到請求的起點用 Background）。
	ctx := context.Background()

	// ── 9. 用 WithValue 把 request 範圍的值「疊」上去 ─────────────
	// 每 WithValue 一次就派生一個新 ctx，帶上一組 key/value。可以疊很多層。
	// 慣例上會包成一個小 helper（withRequestID），讓呼叫端不用直接碰 key。
	ctx = withRequestID(ctx, "req-20260701-4417")
	ctx = context.WithValue(ctx, keyUserID, 88001) // 也可直接呼叫，這裡示範原始寫法

	// 模擬請求進 handler；ctx 帶著這些值一路往下傳。
	handleCheckout(ctx)
}

// withRequestID 是慣例做法：把「塞值」包成函式，呼叫端只給值、看不到 key 細節。
func withRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, keyRequestID, id)
}

// requestIDFrom 是配套的「取值」helper：做好型別斷言與「拿不到」的處理，回傳乾淨的結果。
func requestIDFrom(ctx context.Context) string {
	// ctx.Value 回傳 any（interface{}），要用「逗號 ok」型別斷言轉回 string。
	// 若這個 key 根本沒被設過，v 會是 nil、ok 會是 false。
	if id, ok := ctx.Value(keyRequestID).(string); ok {
		return id
	}
	return "unknown" // 拿不到就給預設值，不要 panic
}

func handleCheckout(ctx context.Context) {
	// ── 10. 在深層讀值：任何拿到 ctx 的地方都能讀，不必層層加參數 ──
	// 這就是好處：requestID 不用出現在每個函式的參數列，卻能在最底層的 log 讀到。
	fmt.Printf("[handler] requestID=%s 開始結帳\n", requestIDFrom(ctx))
	writeAuditLog(ctx, "建立訂單")
}

func writeAuditLog(ctx context.Context, action string) {
	reqID := requestIDFrom(ctx)

	// 讀 userID：直接示範原始的型別斷言（回傳 any → 斷言回 int）。
	userID, ok := ctx.Value(keyUserID).(int)
	if !ok {
		userID = 0 // 沒有登入者就當 0
	}

	// 真實世界這行會寫進 log 系統，requestID 讓你能把同一個請求的所有 log 串起來追蹤。
	fmt.Printf("  [audit] req=%s user=%d 動作=%s\n", reqID, userID, action)
}
