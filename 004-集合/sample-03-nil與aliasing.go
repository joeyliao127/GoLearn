//go:build ignore

// sample-03-nil與aliasing｜對應「教學.html › 8~10 節」
// 電商情境：初始化訂單明細/庫存表時的 nil 行為，以及「切一段購物車給別人改，卻改到原本那份」的 aliasing 陷阱。
// 三個必記重點（電商 code 最常踩的雷）：
//
//	① nil slice 可以直接 append（很方便，逐步把明細加進去）。
//	② nil map 一「寫入」就 panic（必須先 make 或用字面值建好才寫）。
//	③ slice aliasing：slice 只是「指向底層陣列的視窗」，兩個 slice 可能共享同一塊底層陣列，
//	   改一邊會動到另一邊——切子清單交出去時特別容易中招。
//
// 跑法：go run 004-集合/sample-03-nil與aliasing.go
package main

import "fmt"

func main() {
	// ── 8. nil slice：可以直接 append（零值就能用）───────────
	// 只宣告不初始化的 slice 是 nil：len=0、cap=0，但 append 完全 OK。
	// 慣例上要「逐步收集」一個清單時，常直接宣告 var xs []T 然後一路 append。
	var lineItems []string // nil slice，還沒放任何明細
	fmt.Printf("nil slice：len=%d cap=%d，是 nil 嗎？%v\n", len(lineItems), cap(lineItems), lineItems == nil)

	lineItems = append(lineItems, "機械鍵盤") // 對 nil slice append 完全合法
	lineItems = append(lineItems, "滑鼠墊")
	fmt.Printf("append 後：%v（現在還是 nil 嗎？%v）\n", lineItems, lineItems == nil)

	// ── 9. nil map：一「寫入」就 panic！（必須先建好）─────────
	// 只宣告的 map 是 nil。讀 nil map 不會 panic（回零值），但「寫入」nil map 會 runtime panic。
	// 這裡故意示範這個 panic，並用 defer + recover 接住它，讓程式能繼續往下跑。
	demoNilMapWritePanics()

	// 正解：map 一定要先 make（或用字面值建）再寫。
	stock := make(map[string]int) // make 建出一個可用的空 map
	stock["A001"] = 10            // 現在寫入沒問題了
	fmt.Printf("先 make 再寫入就 OK：A001 = %d\n", stock["A001"])

	// ── 10. slice aliasing：兩個 slice 共享同一塊底層陣列 ─────
	// slice 底層是「指向一塊陣列的視窗（指標＋長度＋容量）」。把一個 slice 指派給另一個、
	// 或用 s[low:high] 切一段，得到的是「看同一塊底層陣列的另一個視窗」，不是複製一份資料。
	cart := []string{"鍵盤", "滑鼠", "螢幕", "耳機"}

	// 切出「前兩項」交給促銷模組處理——它以為是自己的一份，其實跟 cart 共用底層陣列。
	promo := cart[0:2] // promo 看到 ["鍵盤","滑鼠"]，但和 cart 共享底層
	promo[0] = "鍵盤(促銷替換)"
	fmt.Printf("改 promo[0] 竟然也改到了 cart[0]：\n  cart  = %v\n  promo = %v\n", cart, promo)

	// 想要「真正獨立的一份」，用 copy 複製到新 slice（或用 append 到 nil slice）。
	independent := make([]string, len(cart))
	copy(independent, cart) // 把 cart 的內容整份複製過來
	independent[0] = "獨立副本改這個"
	fmt.Printf("用 copy 後，改 independent 不影響 cart：\n  cart        = %v\n  independent = %v\n", cart, independent)
}

// demoNilMapWritePanics 故意寫入 nil map 觸發 panic，用 recover 接住，示範第 9 節。
func demoNilMapWritePanics() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("寫入 nil map 觸發 panic（已用 recover 接住）：%v\n", r)
		}
	}()

	var priceTable map[string]int                                // nil map，還沒 make
	fmt.Printf("讀 nil map 不會 panic，回零值：%d\n", priceTable["any"]) // 讀 OK
	priceTable["A001"] = 100                                     // ← 寫入 nil map，這行會 panic
	fmt.Println("這行不會被執行到")                                      // panic 後直接跳到上面的 defer
}
