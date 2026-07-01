//go:build ignore

// sample-02-map｜對應「教學.html › 5~7 節」
// 電商情境：用 map[string]int 當「商品 SKU → 庫存數量」的字典，讀寫庫存、判斷 SKU 是否存在。
// 重點：map 是 Go 的鍵值對字典（其他語言的 dict / hash / HashMap），
// 電商 domain 裡「以 ID 查東西」的地方（庫存表、價目表、購物車數量）幾乎都是它。
// 特別學會 v, ok := m[k] 這個「查存在」慣例——這是 map 最關鍵的 idiom。
// 跑法：go run 004-集合/sample-02-map.go
package main

import "fmt"

func main() {
	// ── 5. 建 map（字面值）＋ 讀寫 ───────────────────────────
	// map[鍵型別]值型別；這裡鍵是 SKU 字串、值是庫存數量 int。
	stock := map[string]int{
		"A001": 12, // 機械鍵盤庫存 12
		"A002": 0,  // 滑鼠墊：帳面存在，但庫存是 0（注意：這跟「查無此商品」不同）
		"A003": 7,  // USB-C 線庫存 7
	}

	// 讀：用 key 取值。
	fmt.Printf("A001 庫存：%d\n", stock["A001"])

	// 寫：出貨扣庫存、進貨補庫存，直接對 key 賦值。
	stock["A001"] -= 3 // 賣出 3 支鍵盤
	stock["A003"] += 5 // 補 5 條線
	fmt.Printf("扣款出貨後 A001 庫存：%d、A003 庫存：%d\n", stock["A001"], stock["A003"])

	// 新增：對「不存在的 key」賦值，就是新增一筆。
	stock["A009"] = 20 // 上架新商品
	fmt.Printf("目前品項數（len）：%d\n", len(stock))

	// ── 6. 「零值陷阱」與 v, ok := 判存在（map 的靈魂 idiom）──
	// 讀一個「不存在的 key」不會報錯，會回傳「值型別的零值」（int 的零值是 0）。
	// 所以單看 stock["X999"]==0 沒辦法分辨「庫存剛好是 0」還是「根本沒這個商品」。
	fmt.Printf("查無此商品 X999，直接讀會得到零值：%d（分不出是缺貨還是不存在）\n", stock["X999"])

	// 正解：用「逗號 ok」寫法。ok 是 bool，告訴你這個 key「到底存不存在」。
	if qty, ok := stock["A002"]; ok {
		fmt.Printf("A002 存在，庫存 %d（帳面有此商品，但目前缺貨）\n", qty)
	} else {
		fmt.Println("A002 不存在")
	}
	if qty, ok := stock["X999"]; ok {
		fmt.Printf("X999 存在，庫存 %d\n", qty)
	} else {
		fmt.Printf("X999 不存在（qty 是零值 %d，但要看 ok=false 才是判斷依據）\n", qty)
	}

	// ── 7. delete 刪 key ＋ for range 走訪 map ───────────────
	// delete(m, key)：下架商品就從庫存表刪掉。刪一個不存在的 key 不會報錯（安全）。
	delete(stock, "A002")
	fmt.Printf("下架 A002 後品項數：%d\n", len(stock))

	// 走訪 map 也用 for range，每輪給 (鍵, 值)。
	// ⚠️ map 的走訪順序「不保證、每次可能不同」——需要固定順序得自己排（見教學小抄）。
	fmt.Println("目前庫存清單（順序不保證）：")
	for sku, qty := range stock {
		fmt.Printf("  %s → %d\n", sku, qty)
	}
}
