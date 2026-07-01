//go:build ignore

// sample-02-逾時與取消｜對應「教學.html › 4~7 節」
// 電商情境：結帳時要呼叫「外部金流閘道」授權刷卡。這種外部呼叫可能很慢，
// 我們不能無限等——給它一個逾時（例如 200ms），超過就放棄、回錯誤給顧客，別卡住整條請求。
// 重點：
//  1. context.WithTimeout 產生「會自己逾時的 ctx」＋一定要 defer cancel()。
//  2. 慢操作用 select 同時等「做完」和「ctx.Done()（被取消/逾時）」，誰先到聽誰的。
//  3. ctx.Err() 會告訴你是 DeadlineExceeded（逾時）還是 Canceled（被主動取消）。
//  4. WithCancel 讓你「主動喊卡」——例如其中一個子任務失敗，就取消其餘還在跑的。
//
// 跑法：go run 011-context/sample-02-逾時與取消.go
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	// ── 4. WithTimeout：逾時比操作短 → 逾時獲勝（顧客等不起）──────
	// 給 200ms 逾時，但金流授權要花 500ms → 逾時先到，操作被放棄。
	fmt.Println("[案例 A] 逾時 200ms，但授權要 500ms")
	chargeWithTimeout(200*time.Millisecond, 500*time.Millisecond)

	fmt.Println("----")

	// ── 5. WithTimeout：操作比逾時快 → 正常完成 ──────────────────
	// 一樣 200ms 逾時，但這次授權只花 50ms → 在期限內完成，拿到結果。
	fmt.Println("[案例 B] 逾時 200ms，授權只要 50ms")
	chargeWithTimeout(200*time.Millisecond, 50*time.Millisecond)

	fmt.Println("----")

	// ── 7. WithCancel：主動喊卡（不是等逾時，是我們決定不要了）────
	fmt.Println("[案例 C] 主動取消：風控判定可疑，喊停授權")
	chargeWithManualCancel()
}

// chargeWithTimeout 示範第 4、5、6 節：建立會逾時的 ctx，用 select 等「授權完成」或「逾時」。
func chargeWithTimeout(timeout, workTime time.Duration) {
	// 從一個父 ctx（這裡用 Background 當根）派生出「timeout 後會自動取消」的子 ctx。
	// 回傳兩個值：ctx（帶逾時）＋ cancel（取消函式）。
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	// ── 6. 黃金慣例：拿到 cancel 立刻 defer cancel() ─────────────
	// 就算逾時會自己取消，也「一定要」呼叫 cancel 釋放它綁的計時器/資源，否則會洩漏。
	// 放在 defer 最保險：不管這個函式從哪條路徑 return，cancel 都會被呼叫。
	defer cancel()

	// 把「授權」丟到背景 goroutine 去做，做完把結果送進 channel。
	// （goroutine / channel 是主題 010 的東西，這裡當工具用：模擬一個「會花時間的外部呼叫」。）
	done := make(chan string, 1) // buffered，避免逾時後這個 goroutine 卡在送值而洩漏
	go func() {
		time.Sleep(workTime) // 模擬外部金流授權的耗時
		done <- "授權碼 AUTH-8891"
	}()

	// select：同時等兩件事，誰先發生就走誰。
	select {
	case <-ctx.Done():
		// ctx 被取消（這裡是逾時）。ctx.Err() 說明原因。
		// 逾時的 err 會是 context.DeadlineExceeded，可用 errors.Is 判斷。
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			fmt.Println("  逾時放棄！回覆顧客：金流忙碌，請稍後再試（err =", ctx.Err(), "）")
		} else {
			fmt.Println("  被取消：", ctx.Err())
		}
	case auth := <-done:
		// 授權在期限內完成。
		fmt.Println("  授權成功：", auth)
	}
}

// chargeWithManualCancel 示範第 7 節：用 WithCancel 主動喊停，ctx.Err() 會是 Canceled。
func chargeWithManualCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 一樣：無論如何都要呼叫 cancel

	done := make(chan string, 1)
	go func() {
		time.Sleep(300 * time.Millisecond) // 授權要 300ms
		done <- "授權碼 AUTH-0001"
	}()

	// 模擬：風控在 50ms 時判定這筆可疑，主動喊卡。
	go func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("  風控：判定可疑，呼叫 cancel() 取消授權")
		cancel() // 主動取消 → ctx.Done() 會關閉、ctx.Err() 變 context.Canceled
	}()

	select {
	case <-ctx.Done():
		fmt.Println("  已中止授權，原因：", ctx.Err()) // context.Canceled
	case auth := <-done:
		fmt.Println("  授權成功：", auth)
	}
}
