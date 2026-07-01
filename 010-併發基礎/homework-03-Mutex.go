//go:build ignore

// homework-03-Mutex｜搭配 sample-03-Mutex
// 目標（≤5 分鐘）：用 sync.Mutex 保護一個共享計數器，讓 200 個並行 +1 精準加到 200。
// 電商情境：後台同時有 200 筆訂單完成，每筆都要把「今日已完成訂單數」+1。
// 這是多個 goroutine 同時改「同一個變數」，不上鎖就會漏加（data race）。
//
// TODO 1：宣告 var mu sync.Mutex、counter := 0，再一個 var wg sync.WaitGroup
// TODO 2：開 200 個 goroutine（每個先 wg.Add(1)、裡面 defer wg.Done()），用 mu.Lock() → counter++ → mu.Unlock() 保護（慣例：Lock 後立刻 defer mu.Unlock()）
// TODO 3：wg.Wait() 後印「今日已完成訂單數：X」
//
// 完成後跑：go run 010-併發基礎/homework-03-Mutex.go
// 也可加 -race 檢查：go run -race 010-併發基礎/homework-03-Mutex.go（上鎖後應顯示無資料競爭）
// 驗收：印出「今日已完成訂單數：200」（少了鎖幾乎一定 <200；-race 也會報 DATA RACE）
package main

import (
	"fmt"
	"sync"
)

func main() {
	const orders = 200

	// 下面這行只是讓半成品能編譯；宣告了自己的 mu 之後就把這行刪掉。
	var _ sync.Mutex

	// 在這裡開始寫 ↓
	fmt.Println("TODO: 還沒開始，把 TODO 1~3 完成後刪掉這行（目標讓計數精準到", orders, "）")
}
