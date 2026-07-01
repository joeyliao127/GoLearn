//go:build ignore

// sample-04-panic與recover｜對應「教學.html › 10~12 節」
// 電商情境：批次結算一批訂單。其中一筆內部資料壞掉（除以 0 的例子模擬 bug），
// 我們不希望「一筆炸掉就整個批次程式掛掉」，於是用 defer + recover 攔住 panic，
// 把它「降級」成一個普通 error 繼續處理下一筆。
// 重點：① panic 用於「真正異常/程式不變式被打破」，不是拿來當一般錯誤流程；
// ② recover 只有寫在 defer 的函式裡才有效，能攔住 panic、讓程式不崩；
// ③ 平常「可預期的失敗」（缺貨、餘額不足）請回 error，別 panic（見 sample-01~03）。
// 跑法：go run 008-錯誤處理/sample-04-panic與recover.go
package main

import (
	"errors"
	"fmt"
)

// ── 10. panic 的時機：真正的異常，不是一般錯誤 ──────────────
// divShare 把金額平均拆給 n 個人。n==0 在這個系統裡「本不該發生」（呼叫端的 bug），
// 用 panic 表達「這是不變式被打破，不是正常業務分支」。
// （對比：缺貨是可預期的，那種要回 error；panic 留給「這根本不該發生」。）
func divShare(amount, n int) int {
	if n == 0 {
		panic("divShare: 人數不可為 0（呼叫端邏輯有誤）")
	}
	return amount / n
}

// ── 11. recover 只在 defer 裡有效：把 panic 攔成 error ────────
// settleOne 用「具名回傳值 err」+ defer 攔截 panic：
// 一旦內部 panic，defer 裡的 recover() 會拿到 panic 值、阻止程式崩潰，
// 我們再把它塞回 err 回傳，對外看起來就只是「這筆失敗了」。
func settleOne(orderID string, amount, n int) (err error) {
	defer func() {
		// recover() 在「正常結束」時回 nil；發生 panic 時回 panic 的值。
		if r := recover(); r != nil {
			// 把 panic 降級成一般 error（%v 印出 panic 值）回傳出去。
			err = fmt.Errorf("結算訂單 %s 時發生嚴重錯誤: %v", orderID, r)
		}
	}()

	share := divShare(amount, n) // n==0 時這裡會 panic，被上面的 defer 接住
	fmt.Printf("訂單 %s：每人分攤 %d 元\n", orderID, share)
	return nil
}

func main() {
	// 一批訂單，故意讓中間一筆 n==0 觸發 panic。
	type job struct {
		id     string
		amount int
		n      int
	}
	jobs := []job{
		{"ORD-1", 1000, 4}, // 正常
		{"ORD-2", 900, 0},  // 內部異常 → 會 panic，但被 recover 攔住降級成 error
		{"ORD-3", 600, 3},  // 前一筆沒把程式弄垮，這筆照跑
	}

	for _, j := range jobs {
		// ── 12. 收尾慣例：靠 defer+recover，一筆炸掉不拖垮整批 ──
		if err := settleOne(j.id, j.amount, j.n); err != nil {
			fmt.Println("這筆略過：", err)
			continue // 降級成 error 後，正常用 if err != nil 處理，批次繼續
		}
	}

	fmt.Println("批次結束：程式沒有崩潰（panic 被 recover 接住了）")

	// 最後對照一下 sample-01~03 的正常做法：可預期的失敗回 error，不要用 panic。
	if err := errors.New("（示意）缺貨這種可預期的失敗 → 回 error，不 panic"); err != nil {
		fmt.Println(err)
	}
}
