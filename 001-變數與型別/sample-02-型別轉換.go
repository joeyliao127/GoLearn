//go:build ignore

// sample-02-型別轉換｜對應「教學.html › 6~8 節」
// 電商情境：算一筆訂單的折扣後金額與每件實付，會用到 int↔float64 轉換。
// 跑法：go run 001-變數與型別/sample-02-型別轉換.go
package main

import "fmt"

func main() {
	// ── 6. 基本型別 ───────────────────────────────────────
	quantity := 3        // int：整數（數量）
	unitPrice := 590     // int：單價（元，用整數存錢）
	discountRate := 0.85 // float64：折扣（打 85 折）

	// ── 7. 運算符：訂單小計 = 單價 × 數量 ──────────────────
	subtotal := unitPrice * quantity // int * int = int
	fmt.Printf("小計：%d 元\n", subtotal)

	// ── 8. 型別轉換：Go 不會自動混算不同型別 ──────────────
	// subtotal 是 int、discountRate 是 float64，「不能直接相乘」，
	// 必須先把 int 顯式轉成 float64：float64(subtotal)
	total := float64(subtotal) * discountRate // float64 * float64 = float64
	fmt.Printf("折扣後：%.1f 元\n", total)         // %.1f：印到小數點後 1 位

	// float64 → int 會「直接截斷小數」（不是四捨五入）
	fmt.Printf("折扣後(取整/截斷)：%d 元\n", int(total))

	// 每件實付：total ÷ 數量。total 已是 float64，quantity 是 int，要先轉：
	avgPerItem := total / float64(quantity)
	fmt.Printf("每件實付：%.2f 元\n", avgPerItem)

	// ── 整數除法陷阱：int / int 會無條件捨去小數 ───────────
	// 例：100 元均分給 3 個人
	fmt.Printf("陷阱：100/3（int）= %d，但 float64 = %.2f\n",
		100/3, float64(100)/float64(3))

	// 取餘數 %：常用於整除判斷、分頁、奇偶（Printf 裡印 % 要寫 %%）
	fmt.Printf("100 %% 3 的餘數 = %d\n", 100%3)
}
