//go:build ignore

// sample-01-多回傳值｜對應「教學.html › 1~4 節」
// 電商情境：結帳時把一筆訂單「拆帳」——同時算出「小計」與「稅額」兩個結果；
// 再示範 Go 最核心的慣例：函式回傳 (值, error)，呼叫端立刻 if err != nil 檢查。
// 重點：① 一次回傳多個值；② 具名回傳值 (subtotal, tax int)；
// ③ (值, error) 是 Go 的標準慣例；④ naked return 少用（可讀性差）。
// 跑法：go run 003-函式/sample-01-多回傳值.go
package main

import (
	"errors"
	"fmt"
)

func main() {
	// ── 1. 多回傳值：一次接住兩個結果 ────────────────────────
	// Go 函式可以回傳不只一個值，呼叫端用「多重指派」一次接住。
	subtotal, tax := priceBreakdown(1000, 3) // 單價 1000、數量 3
	fmt.Printf("小計 %d 元，稅額 %d 元，應付 %d 元\n", subtotal, tax, subtotal+tax)

	// 用不到其中某個回傳值時，用 _ 明確丟掉（Go 不准宣告了卻不用的變數）。
	var onlyTax int
	_, onlyTax = priceBreakdown(500, 2)
	fmt.Printf("只關心稅額：%d 元\n", onlyTax)

	// ── 3~4. (值, error) 慣例：成功給值、失敗給 error ─────────
	// 這是 Go 最重要的慣例：把「可能失敗」的操作寫成回傳 (結果, error)，
	// 呼叫端「馬上」用 if err != nil 檢查，先處理錯誤、再往下走（早 return）。
	unitPrice, err := applyCoupon(1000, "SAVE100")
	if err != nil {
		fmt.Println("套用折扣碼失敗：", err)
	} else {
		fmt.Printf("折扣後單價：%d 元\n", unitPrice)
	}

	// 故意給一個無效的折扣碼，看錯誤路徑
	if _, err := applyCoupon(1000, "FAKE"); err != nil {
		fmt.Println("套用折扣碼失敗：", err) // 慣例：if 帶初始化，err 只活在這個 if 內
	}
}

// ── 2. 具名回傳值：在回傳型別處先幫回傳值取名 ────────────────
// (subtotal, tax int) 等於在函式一開始就宣告好 subtotal、tax（都是 int，零值 0）。
// 好處是「回傳了什麼」寫在簽名上、一看就懂；函式內直接對它們賦值即可。
func priceBreakdown(unitPrice, qty int) (subtotal, tax int) {
	subtotal = unitPrice * qty
	tax = subtotal * 5 / 100 // 稅率 5%；金額用 int 純整數運算，別用 float
	return subtotal, tax     // 慣例：明確寫出要回傳什麼，別用 naked return（見教學 trap）
}

// applyCoupon 示範 (值, error) 慣例：折扣碼有效就回 (折後價, nil)，無效就回 (0, error)。
// 慣例：error 一律放「最後一個」回傳值；成功時 error 為 nil。
func applyCoupon(unitPrice int, code string) (int, error) {
	switch code {
	case "SAVE100":
		return unitPrice - 100, nil // 成功：回結果 + nil
	case "SAVE50":
		return unitPrice - 50, nil
	default:
		// 失敗：回「零值 + 一個描述錯誤的 error」。errors.New 造一個簡單錯誤。
		return 0, errors.New("無效的折扣碼：" + code)
	}
}
