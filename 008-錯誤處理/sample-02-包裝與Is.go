//go:build ignore

// sample-02-包裝與Is｜對應「教學.html › 4~6 節」
// 電商情境：下單前檢查庫存。底層 checkStock 發現缺貨 → 回傳 ErrOutOfStock；
// 上層 placeOrder 把它「往上包一層」加上是哪張訂單、哪個商品的脈絡；
// 最外層再用 errors.Is 判斷「這串錯誤裡到底是不是缺貨」，決定要不要提示補貨。
// 重點：① 用 fmt.Errorf("...: %w", err) 包裝，保留原始錯誤又補上下文；
// ② 包裝後不能再用 == 比對（外皮不一樣了），要用 errors.Is 穿透整條鏈比對；
// ③ errors.Unwrap 可一層層剝開，看見包裝的結構。
// 跑法：go run 008-錯誤處理/sample-02-包裝與Is.go
package main

import (
	"errors"
	"fmt"
)

// sentinel error（哨兵錯誤）：一個「可被比對的固定錯誤值」，慣例名字 ErrXxx。
// sample-03 會專門講它；這裡先用它當「被包裝、被 Is 比對」的對象。
var ErrOutOfStock = errors.New("庫存不足")

// checkStock 是底層：庫存不夠就回傳「原味」的 ErrOutOfStock。
func checkStock(productID string, want int) error {
	stock := map[string]int{"A100": 3, "B200": 0}
	if stock[productID] < want {
		return ErrOutOfStock // 直接回傳 sentinel 本尊
	}
	return nil
}

// ── 4. 用 %w 包裝：保留原錯誤，再補上「是誰、在哪出事」的脈絡 ──
// placeOrder 是上層：它呼叫 checkStock，若出錯就「包一層」加上訂單/商品資訊。
// 關鍵是動詞 %w（wrap）：它把 err 藏進新錯誤裡，之後 errors.Is/As 還能挖出來。
// （對比：若寫成 %v 只會複製「文字」，原始錯誤的「身分」就斷掉了，Is 也就比不到。）
func placeOrder(orderID, productID string, qty int) error {
	if err := checkStock(productID, qty); err != nil {
		return fmt.Errorf("訂單 %s 建立失敗（商品 %s ×%d）: %w", orderID, productID, qty, err)
	}
	return nil
}

func main() {
	// 商品 B200 庫存為 0，會一路缺貨錯誤上來。
	err := placeOrder("ORD-0001", "B200", 2)

	// ── 5. 包裝後的錯誤：文字含「上下文 + 原因」，但別再用 == 比 ──
	fmt.Println("完整錯誤訊息：", err)
	// 這樣比對永遠是 false！因為 err 現在是「被包過的外皮」，不再等於 ErrOutOfStock 本尊。
	fmt.Println("用 == 直接比（會失手）：", err == ErrOutOfStock)

	// ── 6. errors.Is：穿透整條包裝鏈，比對「是不是某個 sentinel」 ──
	// Is 會沿著 %w 一層層 Unwrap，只要鏈上任一層 == ErrOutOfStock 就回 true。
	if errors.Is(err, ErrOutOfStock) {
		fmt.Println("errors.Is 判斷：確實是缺貨 → 顯示『到貨通知我』按鈕")
	}

	// 補充：errors.Unwrap 可以手動剝一層，看見底下包的是誰（實務多半用 Is，不手剝）。
	fmt.Println("剝一層看到的原因：", errors.Unwrap(err))

	// 對照組：商品 A100 有 3 個，買 2 個沒問題，err 會是 nil。
	if err := placeOrder("ORD-0002", "A100", 2); err == nil {
		fmt.Println("ORD-0002 下單成功（庫存足夠）")
	}
}
