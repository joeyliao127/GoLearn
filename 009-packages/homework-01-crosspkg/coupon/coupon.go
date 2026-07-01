// Package coupon｜練習：折價券計算。示範「該匯出的匯出、該藏的藏起來」。
package coupon

// Apply 已經幫你匯出好（大寫），main 會跨包呼叫它。
// TODO：實作「折價後金額 = price - discount，但最低 0（不能變負數）」。
// 提示：用下面那個「未匯出」的 clampZero 把負數夾成 0。
func Apply(price, discount int) int {
	// 先回 0 佔位讓半成品能編譯；完成後改成正確邏輯（記得用 clampZero）。
	return 0
}

// clampZero 小寫未匯出：只有 coupon 內部能用（main 看不到它）。負數夾成 0。
func clampZero(n int) int {
	if n < 0 {
		return 0
	}
	return n
}
