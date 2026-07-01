//go:build ignore

// sample-02-variadic｜對應「教學.html › 5~6 節」
// 電商情境：算購物車金額——商品件數不固定，用「可變參數」一次收任意多個價格加總。
// 重點：① variadic 參數寫成 ...T，函式內它就是一個 []T；② 呼叫時可帶 0、1 或多個；
// ③ 已經有一個 slice 時，用 slice... 把它「攤開」傳進去；④ variadic 只能是最後一個參數。
// 跑法：go run 003-函式/sample-02-variadic.go
package main

import "fmt"

func main() {
	// ── 5. 呼叫 variadic：想傳幾個就傳幾個 ───────────────────
	// sum 的參數是 prices ...int，呼叫時直接把價格一個個列上去即可。
	fmt.Println("三件商品合計：", sum(2490, 290, 190)) // 傳 3 個
	fmt.Println("一件商品合計：", sum(999))            // 傳 1 個
	fmt.Println("空購物車合計：", sum())               // 傳 0 個 → 回 0

	// ── 6. 用 slice... 把既有的 slice 攤開傳進去 ─────────────
	// 常見情況：價格本來就裝在一個 []int 裡，這時不能直接把 slice 當一個值傳，
	// 要在後面加 ...，Go 才會把它「攤平」成一個個引數餵給 variadic 參數。
	cartPrices := []int{1200, 350, 80, 80}
	fmt.Println("購物車（用 slice... 攤開）合計：", sum(cartPrices...))

	// ── 混合固定參數 + variadic：variadic 一定放最後 ─────────
	// 前面可以有普通參數（這裡 customer），可變參數 items 一定是「最後一個」。
	fmt.Println(orderSummary("小美", 2490, 290, 190))
	fmt.Println(orderSummary("阿宏")) // items 傳 0 個也 OK，函式內就是空 slice
}

// sum 把任意多個價格加總。參數 prices ...int 在函式「內部」其實就是一個 []int，
// 所以可以直接 for range 走訪它。金額全用 int，別用 float。
func sum(prices ...int) int {
	total := 0
	for _, p := range prices { // prices 就是 []int
		total += p
	}
	return total
}

// orderSummary 示範「固定參數 + variadic」混用：customer 是普通參數，
// items ...int 是可變參數、且必須放在參數列的最後。
func orderSummary(customer string, items ...int) string {
	return fmt.Sprintf("%s 的訂單：%d 件商品，合計 %d 元", customer, len(items), sum(items...))
}
