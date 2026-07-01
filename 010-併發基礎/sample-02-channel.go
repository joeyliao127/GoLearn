//go:build ignore

// sample-02-channel｜對應「教學.html › 4~6 節」
// 電商情境：一批訂單進來，開幾個「worker」並行處理（算每筆的總價），
// 用 channel 當輸送帶：一條 jobs 送訂單進去、一條 results 把算好的金額收回來。
// 這就是 Go 的核心心法：「用通訊來共享記憶體」—— 資料在 channel 上流動，
// 而不是大家去搶同一塊變數（那才需要鎖，見 sample-03）。
// 重點：① make(chan T) 建 channel、ch <- v 送、v := <-ch 收；
// ② close(ch) + for range ch 收到關閉為止；③ worker pool 的雛形。
// 跑法：go run 010-併發基礎/sample-02-channel.go
package main

import (
	"fmt"
	"sync"
)

// 一筆待處理的訂單（金額用 int 存「元」，別用 float 記錢）。
type Order struct {
	ID    int
	Price int
	Qty   int
}

func main() {
	orders := []Order{
		{ID: 1, Price: 2490, Qty: 1},
		{ID: 2, Price: 290, Qty: 3},
		{ID: 3, Price: 190, Qty: 2},
		{ID: 4, Price: 999, Qty: 1},
	}

	// ── 4. 建 channel：jobs 送訂單、results 收金額 ──────────────────
	// channel 是「有型別」的輸送帶：jobs 只能流 Order、results 只能流 int。
	jobs := make(chan Order)
	results := make(chan int)

	// ── 5. 起 3 個 worker，並行從 jobs 收訂單、把小計送進 results ─────
	const workerCount = 3
	var wg sync.WaitGroup
	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// 送訂單進 jobs；送完 close(jobs) 告訴 worker「沒有更多工作了」。
	// 用一個 goroutine 來送，才不會和下面收 results 互相卡住（deadlock）。
	go func() {
		for _, o := range orders {
			jobs <- o // 送一筆訂單上輸送帶（沒 worker 來收就會在這裡等）
		}
		close(jobs) // 關閉：worker 的 for range 收到「關閉」就會自然結束
	}()

	// 另一個 goroutine 專門等所有 worker 做完，然後 close(results)，
	// 這樣下面的 for range results 才知道「收完了、可以停」。
	go func() {
		wg.Wait()
		close(results)
	}()

	// ── 6. for range 一個 channel：一直收，直到它被 close ────────────
	// 主流程在這裡把每個 worker 回傳的小計加總。收的順序不保證，但總和固定。
	total := 0
	for subtotal := range results {
		total += subtotal
	}
	fmt.Printf("全部訂單處理完畢，合計金額 %d 元\n", total)
}

// worker 從 jobs 一直收訂單處理，把每筆小計送進 results。
// jobs 這種 <-chan 只收、chan<- 只送的「方向型別」是慣例：把用途寫進型別、更安全。
func worker(id int, jobs <-chan Order, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for o := range jobs { // jobs 被 close 後，這個迴圈就結束
		subtotal := o.Price * o.Qty
		fmt.Printf("worker %d 處理訂單 %d：%d 元\n", id, o.ID, subtotal)
		results <- subtotal
	}
}
