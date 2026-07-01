//go:build ignore

// homework-02-逾時與取消｜搭配 sample-02-逾時與取消
// 目標（≤5 分鐘）：幫一個「寄出貨通知」的慢操作加上逾時，超過就放棄。
// 電商情境：出貨後要呼叫「簡訊供應商」發通知。供應商這次很慢（要 300ms），
// 但我們只願意等 100ms —— 超過就放棄、印出逾時訊息，別卡住流程。
//
// 已幫你備好：背景 goroutine 會花 300ms 後把結果送進 done channel（不用改它）。
//
// TODO 1：用 context.WithTimeout 從 parent 建「100ms 逾時」的 ctx 與 cancel。
// 提示：ctx, cancel := context.WithTimeout(parent, 100*time.Millisecond)
// TODO 2：緊接著 defer cancel()（黃金慣例：拿到 cancel 就 defer）。
// TODO 3：用 select 同時等兩件事：case <-ctx.Done() 印「逾時放棄寄送通知」；
// case msg := <-done 印「通知已寄出：<msg>」。
//
// 完成後跑：go run 011-context/homework-02-逾時與取消.go
// 驗收：因為 100ms < 300ms，逾時先到 → 印出「逾時放棄寄送通知」。
// 試把逾時改成 500ms（>300ms），就會改印「通知已寄出：SMS-OK」。
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 背景把「寄簡訊」跑起來，300ms 後回結果（buffered channel，逾時了也不會卡住這個 goroutine）。
	done := make(chan string, 1)
	go func() {
		time.Sleep(300 * time.Millisecond)
		done <- "SMS-OK"
	}()

	// 根 ctx 先幫你備好，TODO 1 從它派生出帶逾時的 ctx。
	parent := context.Background()

	// 下面兩行 _ = ... 只是讓半成品能編譯；用到 parent / done 之後就把對應那行刪掉。
	_ = parent
	_ = done

	// 在這裡開始寫 ↓（TODO 1~3）
	fmt.Println("TODO: 還沒開始，把 TODO 1~3 完成後刪掉這行")
}
