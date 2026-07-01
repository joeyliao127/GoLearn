//go:build ignore

// homework-03-傳值｜搭配 sample-03-傳值
// 目標（≤5 分鐘）：把一個 request 範圍的值（requestID）放進 ctx，再從深層讀出來。
// 電商情境：請求進來時帶一個 requestID，log 需要它來追蹤同一個請求的所有動作。
//
// 已幫你備好：未匯出的 key 型別 ctxKey 與常數 keyRequestID（這是把值放進 ctx 的正確 key 做法，不要改）。
//
// TODO 1：在 main 用 context.WithValue(ctx, keyRequestID, "req-777") 把 requestID 疊進 ctx，再把新 ctx 傳給 logAction(ctx, "付款")。
// TODO 2：完成 logAction——用 ctx.Value(keyRequestID) 取值、做「逗號 ok」型別斷言轉回 string（取不到用 "unknown"），最後印「[log] req=<id> 動作=<action>」。
//
// 完成後跑：go run 011-context/homework-03-傳值.go
// 驗收：印出「[log] req=req-777 動作=付款」；若把 TODO 1 的 WithValue 那行拿掉，會印成 req=unknown
package main

import (
	"context"
	"fmt"
)

// 未匯出的 key 型別（照 sample-03 的慣例，別用字串當 key）。這段不用改。
type ctxKey int

const keyRequestID ctxKey = iota

func logAction(ctx context.Context, action string) {
	// TODO 2：從 ctx 取 keyRequestID 的值、型別斷言回 string（取不到用 "unknown"），
	//         再印「[log] req=<id> 動作=<action>」。
	_ = ctx    // 開始用 ctx 後刪掉這行
	_ = action // 開始用 action 後刪掉這行
	fmt.Println("TODO: 完成 logAction（TODO 2）")
}

func main() {
	ctx := context.Background()

	// 下面這行 _ = ... 只是讓半成品能編譯；TODO 1 用到 ctx 之後就把這行刪掉。
	_ = ctx

	// 在這裡開始寫 ↓（TODO 1：WithValue 疊上 requestID，再呼叫 logAction）
	fmt.Println("TODO: 還沒開始，把 TODO 1~2 完成後刪掉這行")
}
