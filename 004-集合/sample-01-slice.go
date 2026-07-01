//go:build ignore

// sample-01-slice｜對應「教學.html › 1~4 節」
// 電商情境：把購物車裡的商品用 []Product（slice）裝起來，走訪、加購（append）、加總金額。
// 重點：slice 是 Go 最常用的「可變長度清單」，電商 domain 的訂單明細、商品列表幾乎都是它。
// 順便看懂 array 與 slice 的差別（為什麼實務幾乎只用 slice）、以及 len / cap / append 的行為。
// 跑法：go run 004-集合/sample-01-slice.go
package main

import "fmt"

// 購物車裡的一項商品（金額用 int 存「元」，別用 float，浮點數算錢會有誤差）。
type Product struct {
	Name  string
	Price int
}

func main() {
	// ── 1. 用「字面值」建一個 slice：最常見的起手式 ───────────
	// []Product{...} 就是 slice 字面值；一開始塞 3 項商品進購物車。
	// slice 不用先講長度（跟下面第 3 節的 array 不同），要幾個就給幾個。
	cart := []Product{
		{Name: "機械鍵盤", Price: 2490},
		{Name: "滑鼠墊", Price: 290},
		{Name: "USB-C 線", Price: 190},
	}
	fmt.Printf("目前購物車有 %d 項（len=%d, cap=%d）\n", len(cart), len(cart), cap(cart))

	// ── 2. append：往 slice 尾巴加東西（slice 的靈魂）─────────
	// append 慣例：一定要「把回傳值指回原變數」 cart = append(cart, ...)。
	// 因為容量不夠時 append 會配一塊更大的底層陣列、回傳「新的 slice」，
	// 沒接回去的話就白加了。可以一次 append 多個。
	cart = append(cart, Product{Name: "手機支架", Price: 350})
	cart = append(cart, Product{Name: "清潔布", Price: 90}, Product{Name: "束線帶", Price: 60})
	fmt.Printf("加購後有 %d 項\n", len(cart))

	// ── 3. 走訪 slice 加總金額：for range 是標配 ──────────────
	// range 每輪回傳 (索引, 元素副本)；用不到索引就用 _ 忽略。
	total := 0
	for _, p := range cart {
		total += p.Price
		fmt.Printf("  - %s：%d 元\n", p.Name, p.Price)
	}
	fmt.Printf("購物車小計：%d 元\n", total)

	// ── 4. 用索引存取 + slicing 取子清單 ─────────────────────
	// 索引從 0 開始；cart[0] 是第一項。
	fmt.Printf("第一項是：%s\n", cart[0].Name)

	// slice 表達式 cart[low:high]：取「low 到 high-1」這一段（含頭不含尾）。
	// 這裡取前 3 項當「本次推薦」。省略 low 預設 0、省略 high 預設 len。
	top3 := cart[:3]
	fmt.Printf("前 3 項推薦共 %d 項：", len(top3))
	for i, p := range top3 {
		if i > 0 {
			fmt.Print("、")
		}
		fmt.Print(p.Name)
	}
	fmt.Println()

	// ── 補充：array（固定長度）長什麼樣，對照一下 ────────────
	// array 要在型別就寫死長度 [3]int，長度是型別的一部分，不能變。
	// 實務上幾乎都用 slice，array 少見——這裡只是讓你認得它、知道差在哪。
	var quarterlySales [3]int // 一季三個月的銷售額，長度固定為 3
	quarterlySales[0] = 12000
	quarterlySales[1] = 15000
	quarterlySales[2] = 9000
	fmt.Printf("array 範例（固定 3 格）：%v，len=%d\n", quarterlySales, len(quarterlySales))
}
