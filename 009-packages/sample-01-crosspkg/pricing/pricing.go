// Package pricing 示範「一個資料夾 = 一個 package」與大小寫匯出規則。
// 電商情境：把「算含稅金額」的邏輯獨立成一個 package，給別處重用。
package pricing

// Total 大寫開頭 → 匯出（exported）：別的 package 可以呼叫。
// 算「未稅小計 + 稅」= 含稅總額。
func Total(unitPrice, qty int) int {
	subtotal := unitPrice * qty
	return subtotal + applyTax(subtotal)
}

// applyTax 小寫開頭 → 未匯出（unexported）：只有 pricing 內部能用。
// 別的 package 想呼叫會編譯失敗（cannot refer to unexported name）。
func applyTax(amount int) int {
	const taxRate = 5 // 5%
	return amount * taxRate / 100
}
