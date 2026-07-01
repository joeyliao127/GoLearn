//go:build ignore

// sample-03-for迴圈｜對應「教學.html › 7~9 節」
// 電商情境：走訪購物車，加總「總金額」與「總件數」。
// 重點：Go 只有 for 一個迴圈關鍵字，沒有 while / do-while，一個 for 就能寫出
// 「標準三段式」「當作 while」「無限迴圈」三種形態，而走訪 slice/map 幾乎都用 for range。
// 跑法：go run 002-流程控制/sample-03-for迴圈.go
package main

import "fmt"

// 購物車裡的一項商品（金額用 int 存「元」，別用 float）。
type CartItem struct {
	Name  string
	Price int
	Qty   int
}

func main() {
	cart := []CartItem{
		{Name: "機械鍵盤", Price: 2490, Qty: 1},
		{Name: "滑鼠墊", Price: 290, Qty: 2},
		{Name: "USB-C 線", Price: 190, Qty: 3},
	}

	// ── 7. for range：走訪 slice，最常用的形態 ───────────────
	// range 每輪回傳 (索引, 元素副本)。用不到索引就用 _ 忽略。
	totalAmount := 0
	totalQty := 0
	for _, item := range cart {
		subtotal := item.Price * item.Qty
		totalAmount += subtotal
		totalQty += item.Qty
		fmt.Printf("- %s ×%d＝%d 元\n", item.Name, item.Qty, subtotal)
	}
	fmt.Printf("購物車共 %d 件，合計 %d 元\n", totalQty, totalAmount)

	// ── 8. for 三段式：init; condition; post（傳統計數迴圈）─────
	// 想要「第幾項」時就用索引跑；這裡示範列出前 N 項當作推薦。
	fmt.Println("前 2 項推薦：")
	for i := 0; i < len(cart) && i < 2; i++ {
		fmt.Printf("  %d. %s\n", i+1, cart[i].Name)
	}

	// ── 9. for 當 while 用：只留「條件」那一段 ────────────────
	// Go 沒有 while；把三段式只保留中間的條件，就是 while。
	// 情境：湊免運，一直加購直到金額達門檻。
	const freeShippingThreshold = 5000
	amount := totalAmount
	added := 0
	for amount < freeShippingThreshold { // ← 這行就等於其他語言的 while (amount < ...)
		amount += 200 // 每次加購一件 200 元的小東西
		added++
	}
	fmt.Printf("再加購 %d 件（共 +%d 元）湊到 %d 元免運\n", added, amount-totalAmount, amount)

	// ── 補充：無限迴圈 + break / continue ────────────────────
	// 連條件都不寫就是 for {}（無限迴圈），靠 break 跳出、continue 跳過本輪。
	// 情境：從候選折扣碼裡找出第一個「非空」的來套用。
	codes := []string{"", "", "SAVE50", "WELCOME"}
	i := 0
	for {
		if i >= len(codes) {
			fmt.Println("沒有可用折扣碼")
			break
		}
		if codes[i] == "" {
			i++
			continue // 空字串跳過，換下一個
		}
		fmt.Printf("套用折扣碼：%s\n", codes[i])
		break // 找到就跳出
	}
}
