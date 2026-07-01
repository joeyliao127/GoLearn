//go:build ignore

// sample-01-if判斷｜對應「教學.html › 1~3 節」
// 電商情境：結帳時依「訂單金額」決定運費——滿 999 免運，否則收 60。
// 順便看懂 if / else if / else、比較與邏輯運算子、以及 Go 特有的「if 帶初始化語句」
// 寫法（在電商 code 常跟 err 一起出現）。
// 跑法：go run 002-流程控制/sample-01-if判斷.go
package main

import "fmt"

func main() {
	// ── 1. if / else if / else：最基本的分支 ─────────────────
	// 注意 Go 慣例：條件「不加小括號」，但 { } 一定要有，且 { 跟在同一行。
	orderAmount := 850 // 這筆訂單金額（元，用 int 存錢，別用 float）

	const freeShippingThreshold = 999 // 免運門檻
	const shippingFee = 60            // 未達門檻的運費

	var fee int
	if orderAmount >= freeShippingThreshold {
		fee = 0 // 滿額免運
	} else {
		fee = shippingFee
	}
	fmt.Printf("訂單金額 %d 元 → 運費 %d 元\n", orderAmount, fee)

	// ── 2. 多段 else if：把金額分成幾個級距給不同回饋 ─────────
	// 由上往下逐條比對，命中一條就結束，不會再往下比。
	if orderAmount >= 3000 {
		fmt.Println("回饋：滿 3000，送 100 元購物金")
	} else if orderAmount >= 1500 {
		fmt.Println("回饋：滿 1500，送 50 元購物金")
	} else if orderAmount >= 800 {
		fmt.Println("回饋：滿 800，送 9 折優惠券")
	} else {
		fmt.Println("回饋：未達門檻，加油再買一點！")
	}

	// ── 3. 比較 / 邏輯運算子，與「if 帶初始化語句」 ───────────
	// 比較：== != > < >= <=；邏輯：&& || !（&& 和 || 有短路特性）
	isVIP := true
	if orderAmount >= 500 && isVIP {
		fmt.Println("VIP 且滿 500：本單額外 95 折")
	}

	// Go 特有寫法：if 條件前可放一個「初始化語句」，用 ; 隔開。
	// 這裡宣告的 remain 只活在這個 if / else if / else 內（scope 更小、更安全）。
	// 這正是電商 code 裡到處可見的 `if err := doSomething(); err != nil { ... }` 的骨架。
	if remain := freeShippingThreshold - orderAmount; remain > 0 {
		fmt.Printf("再買 %d 元就免運\n", remain)
	} else {
		fmt.Println("已達免運門檻")
	}
	// 這裡再用 remain 會編譯失敗：它的 scope 只在上面那組 if 裡。
	// fmt.Println(remain) // ❌ undefined: remain
}
