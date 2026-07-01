// Package discount｜練習搭配：這支程式「已經幫你寫好、可運作」，你的任務是幫它補測試。
// 電商情境：算打折後價格。percent 是折扣百分比（10 代表打 9 折）。
package discount

// Apply 回打折後價格：price * (100 - percent) / 100。
// 例：Apply(1000, 10) = 900、Apply(1000, 100) = 0（打到 0 折）、Apply(1000, 0) = 1000。
func Apply(price, percent int) int {
	return price * (100 - percent) / 100
}
