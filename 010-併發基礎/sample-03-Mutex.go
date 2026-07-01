//go:build ignore

// sample-03-Mutex｜對應「教學.html › 7~9 節」
// 電商情境：熱賣商品同時湧入大量下單，很多 goroutine 會「同時扣同一個庫存數字」。
// 如果不保護，兩個 goroutine 可能同時讀到 stock=100、各自算成 99、各自寫回 99，
// 結果賣了兩件卻只扣一件 —— 這叫 data race（資料競爭），會超賣。
// 正解：用 sync.Mutex 把「讀-改-寫」這段圍起來，同一時間只讓一個 goroutine 進去。
// 重點：① Mutex 保護共享狀態；② Lock 後立刻 defer Unlock（成對、不會忘）；
// ③ 有了鎖，100 個並行扣款的結果才會精準是 100（用 go run -race 可驗證無競爭）。
// 跑法：go run 010-併發基礎/sample-03-Mutex.go
// 檢查競爭：go run -race 010-併發基礎/sample-03-Mutex.go（本範例上鎖，應顯示無資料競爭）
package main

import (
	"fmt"
	"sync"
)

// Inventory 是被多個 goroutine 共享的庫存狀態。
// 慣例：把「要保護的資料」和「保護它的鎖」放在同一個 struct，一看就知道誰保護誰。
type Inventory struct {
	mu    sync.Mutex // 保護底下的 stock；零值就能用，不用初始化
	stock int
}

// ── 7. Deduct：把「讀-改-寫」用 Lock/Unlock 圍成不可分割的一段 ─────
func (inv *Inventory) Deduct(n int) bool {
	// ── 8. Lock 後立刻 defer Unlock：不管中途怎麼 return，都保證會解鎖 ──
	inv.mu.Lock()
	defer inv.mu.Unlock()

	if inv.stock < n {
		return false // 庫存不足，扣不了（defer 仍會幫我們 Unlock）
	}
	inv.stock -= n // 這行「同一時間只有一個 goroutine 在做」，所以不會扣錯
	return true
}

// Stock 讀庫存也要上鎖：只要有別的 goroutine 可能同時「寫」，讀就得受保護。
func (inv *Inventory) Stock() int {
	inv.mu.Lock()
	defer inv.mu.Unlock()
	return inv.stock
}

func main() {
	inv := &Inventory{stock: 100}

	// ── 9. 開 100 個 goroutine，每個扣 1 件，全部並行 ────────────────
	// 沒有鎖的話這裡幾乎一定超賣；有了鎖，最後剛好扣到 0。
	const buyers = 100
	var wg sync.WaitGroup
	for i := 0; i < buyers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			inv.Deduct(1) // 每位買家扣 1 件
		}()
	}
	wg.Wait() // 等 100 位買家都扣完

	// 100 件庫存被 100 個並行扣款各扣 1 件，靠 Mutex 保護，結果精準是 0。
	fmt.Printf("初始庫存 100，%d 位買家各扣 1 件後，剩餘庫存 = %d\n", buyers, inv.Stock())
}
