//go:build ignore

// sample-01-變數宣告｜對應「教學.html › 1~5 節」
// 電商情境：宣告一件「商品」的基本欄位，順便看懂零值、const、scope。
// 跑法：go run 001-變數與型別/sample-01-變數宣告.go
package main

import "fmt"

func main() {
	// ── 1. var 的四種宣告寫法 ─────────────────────────────
	// (a) 先宣告、之後再賦值；賦值前它是「零值」
	var name string
	name = "無線滑鼠"

	// (b) 宣告同時賦值，並「明寫型別」
	var price int = 590

	// (c) 省略型別，讓 Go 從右邊自動推斷（這裡推成 int）
	var stock = 20

	// (d) 批量宣告：一組 var 放進括號，讀起來整齊
	var (
		onSale   bool    = true
		weightKg float64 = 0.08
	)

	// ── 2. := 短變數宣告（函式內最常用）──────────────────
	// 等同「var + 自動型別推斷」，但只能寫在函式「裡面」。
	category := "周邊配件"

	fmt.Printf("商品：%s／分類：%s／售價：%d／庫存：%d／上架：%v／重量：%vkg\n",
		name, category, price, stock, onSale, weightKg)

	// ── 3. 零值 zero value：宣告沒給初值時 Go 給的預設 ─────
	// int→0、float64→0、string→""（空字串）、bool→false
	var (
		i int
		f float64
		s string
		b bool
	)
	fmt.Printf("零值：int=%d float64=%v string=%q bool=%v\n", i, f, s, b)

	// ── 4. const 常數：編譯期就固定，執行期不能改 ─────────
	const taxRate = 0.05 // 營業稅 5%，寫死不會變
	fmt.Printf("稅率常數：%v\n", taxRate)

	// ── 5. scope 作用域：{} 內宣告的變數只活在那個區塊 ─────
	{
		flashDiscount := 100 // 這個變數只在這對大括號內存在
		fmt.Printf("限時折抵：%d 元\n", flashDiscount)
	}
	// 離開區塊後再用 flashDiscount 會編譯失敗（已不在它的 scope）：
	// fmt.Println(flashDiscount) // ❌ undefined: flashDiscount
}
