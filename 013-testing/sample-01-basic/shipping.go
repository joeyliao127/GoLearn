// Package shipping｜對應「教學.html › 群組 A（1~2 節）」
// 電商情境：算運費。滿 1000 元免運（回 0），否則收固定 60 元運費。
// 這是「真 package」示範（不是 //go:build ignore 單檔），測試檔就放在同資料夾、同 package。
package shipping

// Fee 依訂單金額回運費：滿 1000 免運、否則 60。金額一律用 int（不用 float）。
func Fee(orderAmount int) int {
	if orderAmount >= 1000 {
		return 0
	}
	return 60
}
