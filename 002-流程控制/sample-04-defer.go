//go:build ignore

// sample-04-defer｜對應「教學.html › 10~12 節」
// 電商情境：處理一筆訂單時，要「鎖定庫存 → 開資料庫連線」，而不論後面成功或失敗，
// 最後都得「解鎖 + 關連線」清理乾淨。defer 就是為這種「保證會執行的收尾」設計的。
// 重點：① defer 在「函式 return 之際」才執行；② 多個 defer 是後進先出（LIFO）；
// ③ 陷阱：defer 的「參數在寫下 defer 那一刻就求值」，不是等到執行時。
// 跑法：go run 002-流程控制/sample-04-defer.go
package main

import "fmt"

func main() {
	// ── 10. defer 的時機：函式結束才執行（模擬資源清理）───────
	processOrder(1001)

	fmt.Println("----")

	// ── 12. 陷阱：defer 參數「當下就求值」 ───────────────────
	deferArgTrap()
}

// processOrder 示範第 10、11 節：清理慣例 + LIFO。
func processOrder(orderID int) {
	fmt.Printf("[開始] 處理訂單 %d\n", orderID)

	// 慣例：一「取得資源」就立刻在下一行 defer「釋放」，成對出現、不會忘。
	// 這兩行的 defer 會排隊等 processOrder return 時才跑。
	fmt.Println("鎖定庫存")
	defer fmt.Println("解鎖庫存") // 先登記 → 後執行

	fmt.Println("開啟 DB 連線")
	defer fmt.Println("關閉 DB 連線") // 後登記 → 先執行

	// ── 11. LIFO：多個 defer 後進先出（像疊盤子，最後放的最先拿）──
	// 中間就算 return 或 panic，上面登記的 defer 依然會執行，清理不漏。
	fmt.Printf("扣款、寫入訂單 %d…完成\n", orderID)

	fmt.Println("[結束] 函式即將 return，接著才輪到 defer 倒序執行 ↓")
	// return 時的執行順序（LIFO）：先「關閉 DB 連線」，再「解鎖庫存」。
}

// deferArgTrap 示範第 12 節：defer 的參數在「defer 當下」就被求值並凍結。
func deferArgTrap() {
	stage := "下單"
	// 這裡 stage 的值（"下單"）當場被複製進 defer；之後 stage 再變也不影響它。
	defer fmt.Println("defer 讀到的 stage =", stage)

	stage = "付款"
	stage = "出貨"
	fmt.Println("函式結尾時 stage =", stage) // 印 "出貨"
	// 但上面那個 defer 到 return 時仍印 "下單"（參數早在登記時就凍結了）。

	// 想改成「用最後的值」可以改包一層匿名函式（閉包在執行時才讀變數）：
	defer func() { fmt.Println("用閉包才讀到最新 stage =", stage) }()
}
