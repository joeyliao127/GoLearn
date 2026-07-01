//go:build ignore

// sample-03-閉包｜對應「教學.html › 7~9 節」
// 電商情境：做一個「購物車累計器」——每呼叫一次就把商品加進去、回傳目前總額；
// 這個「還記得上次總額」的能力，就是靠閉包「捕捉」了外層的變數。
// 重點：① 函式是「值」，可存進變數、當回傳值；② 閉包「捕捉」外層變數（by reference，改得到也記得住）；
// ③ 各自產生的閉包狀態獨立；④ Go 1.22 起 for 迴圈變數每輪獨立。
// 跑法：go run 003-函式/sample-03-閉包.go
package main

import "fmt"

func main() {
	// ── 7~8. 閉包：捕捉外層變數、且能記住狀態 ────────────────
	// makeCart() 回傳一個函式；這個函式「捕捉」了外層的 total，
	// 所以每次呼叫都在同一個 total 上累加——它「記得」上一次的結果。
	addToCart := makeCart()
	fmt.Println("加 2490 後總額：", addToCart(2490)) // 2490
	fmt.Println("加 290  後總額：", addToCart(290))  // 2780
	fmt.Println("加 190  後總額：", addToCart(190))  // 2970

	// ── 9. 各自產生的閉包，狀態互相獨立 ─────────────────────
	// 再叫一次 makeCart() 會得到「全新的、各自獨立」的 total，
	// 兩台購物車互不干擾——證明每個閉包捕捉的是自己那一份變數。
	cartA := makeCart()
	cartB := makeCart()
	fmt.Println("A 加 100：", cartA(100)) // 100
	fmt.Println("A 加 100：", cartA(100)) // 200
	fmt.Println("B 加 500：", cartB(500)) // 500（不受 A 影響）

	// ── 補充：Go 1.22 起，for 迴圈變數「每輪獨立」──────────────
	// 這裡把每個商品名各自包成一個閉包收集起來。
	// 舊 Go（1.21 以前）所有閉包會共用同一個迴圈變數 name，最後全印成最後一項（經典陷阱）；
	// Go 1.22 起每一輪的 name 是獨立副本，閉包各自捕捉到正確的值。
	names := []string{"鍵盤", "滑鼠", "螢幕"}
	var printers []func()
	for _, name := range names {
		printers = append(printers, func() { fmt.Println("結帳商品：", name) })
	}
	for _, p := range printers {
		p() // 依序印「鍵盤 / 滑鼠 / 螢幕」；Go 1.22+ 不會全印成「螢幕」
	}
}

// makeCart 回傳一個閉包：它捕捉了外層的 total，每次呼叫就把 price 累加進 total 並回傳。
// 這正是「函式產生函式 + 捕捉狀態」的典型用法（工廠函式）。
func makeCart() func(price int) int {
	total := 0 // 被下面回傳的匿名函式捕捉；它會「活得比 makeCart 更久」
	return func(price int) int {
		total += price // 改的是被捕捉的那個 total（by reference），所以會累加
		return total
	}
}
