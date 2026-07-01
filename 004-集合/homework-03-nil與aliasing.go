//go:build ignore

// homework-03-nil與aliasing｜搭配 sample-03-nil與aliasing
// 目標（≤5 分鐘）：避免 slice aliasing——複製出一份「獨立的」購物車，改副本時不能動到原本那份。
// 電商情境：要拿購物車做「試算促銷」，得先複製一份來改，不能污染使用者真正的購物車。
//
// TODO 1：用 make 建一個和 cart 一樣長的 backup（backup := make([]string, len(cart))）
// TODO 2：用內建的 copy(backup, cart) 把 cart 的內容整份複製到 backup
// （關鍵：copy 是「複製資料」，backup 和 cart 從此各走各的、不共享底層陣列）
// TODO 3：把 backup[0] 改成 "促銷替換品"，再用 fmt.Printf 印出 cart 和 backup 對照
//
// 完成後跑：go run 004-集合/homework-03-nil與aliasing.go
// 驗收：cart 仍是 [鍵盤 滑鼠 螢幕]（沒被改到）、backup 是 [促銷替換品 滑鼠 螢幕]
//
//	（若你不小心寫成 backup := cart，改 backup 會連 cart 一起變——那就是 aliasing 陷阱）
package main

import "fmt"

func main() {
	cart := []string{"鍵盤", "滑鼠", "螢幕"}

	// 下面這行 _ = ... 只是讓半成品能編譯；開始用 cart 之後就把這行刪掉。
	_ = cart

	// 在這裡開始寫 ↓
	fmt.Println("TODO: 還沒開始，把 TODO 1~3 完成後刪掉這行")
}
