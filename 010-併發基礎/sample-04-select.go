//go:build ignore

// sample-04-select｜對應「教學.html › 10~12 節」
// 電商情境：要向兩家供應商「同時」詢價，誰先回覆就用誰的報價；而且不能傻等，
// 超過 100 毫秒還沒人回就放棄（避免拖垮結帳）。select 就是為「同時等多個 channel、
// 誰先好就處理誰」而生，配上 time.After 就能做「逾時」。這是 011 context 取消/逾時的前導。
// 重點：① select 同時等多個 case，哪個 channel 先就緒就走哪個（都沒好就阻塞）；
// ② time.After(d) 回傳一個「d 之後才會收到值」的 channel，拿來當 timeout 分支；
// ③ default 讓 select 變「非阻塞」（沒人就緒就立刻走 default）。
// 跑法：go run 010-併發基礎/sample-04-select.go
package main

import (
	"fmt"
	"time"
)

// askSupplier 模擬向某供應商詢價：花 delay 時間後，把報價送進一個新 channel 回傳。
// 每次呼叫各自開一個 goroutine 去問，互不阻塞。
func askSupplier(name string, price int, delay time.Duration) <-chan int {
	ch := make(chan int, 1) // 緩衝 1：就算沒人收，goroutine 送完值也能結束、不外漏
	go func() {
		time.Sleep(delay)
		ch <- price
	}()
	return ch
}

func main() {
	// ── 10. select 二選一：兩家同時詢價，誰先回就用誰 ───────────────
	// A 家 60ms、B 家 90ms，兩個 channel 同時等，select 會選先就緒的 A。
	fmt.Println("情境一：兩家同時詢價，取最快回覆的")
	supplierA := askSupplier("A 供應商", 980, 60*time.Millisecond)
	supplierB := askSupplier("B 供應商", 950, 90*time.Millisecond)

	select {
	case p := <-supplierA:
		fmt.Printf("採用 A 供應商報價：%d 元\n", p)
	case p := <-supplierB:
		fmt.Printf("採用 B 供應商報價：%d 元\n", p)
	}

	// ── 11. select + time.After：加一個逾時分支，避免傻等 ───────────
	// 這次問一家很慢的（150ms），但我們只願意等 100ms；time.After 先就緒 → 走逾時。
	fmt.Println("\n情境二：詢價設 100ms 逾時，供應商太慢就放棄")
	slow := askSupplier("C 供應商", 900, 150*time.Millisecond)

	select {
	case p := <-slow:
		fmt.Printf("採用 C 供應商報價：%d 元\n", p)
	case <-time.After(100 * time.Millisecond):
		// time.After 回傳的 channel 在 100ms 後送出一個值 → 命中這個 case。
		fmt.Println("詢價逾時，改用預設價 999 元")
	}

	// ── 12. select + default：非阻塞地「看一眼」channel 有沒有東西 ───
	// default 讓 select「不阻塞」：如果沒有任何 case 當下就緒，就立刻走 default。
	fmt.Println("\n情境三：非阻塞探看庫存回報 channel")
	stockReport := make(chan string) // 沒人往裡面送
	select {
	case msg := <-stockReport:
		fmt.Println("收到庫存回報：", msg)
	default:
		fmt.Println("目前沒有庫存回報，先不等，繼續往下做別的事")
	}
}
