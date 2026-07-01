//go:build ignore

// homework-02-時間｜搭配 sample-02-時間
// 目標（≤5 分鐘）：算「訂單下單到現在幾天」，並算出「下單 +7 天的到期日」。
// 電商情境：後台要顯示這筆訂單「已成立幾天」，還要標出「7 天鑑賞期到期日」。
// （為了讓答案可預期，這裡用固定的 placedAt 與 now，不用 time.Now()。）
//
// TODO 1：用 now.Sub(placedAt) 得到 Duration，再用 .Hours()/24 算出「幾天」（float64），印出「已成立 X 天」（用 %.0f 印整數天）。
// TODO 2：用 placedAt.AddDate(0, 0, 7) 算到期日，用 time.RFC3339 格式印「到期日: ...」。
//
// 提示：對照 sample-02-時間 的第 7、8 節。
// 完成後跑：go run 012-JSON與時間/homework-02-時間.go
// 驗收：已成立 5 天；到期日為 2026-01-17T09:30:00Z
package main

import (
	"fmt"
	"time"
)

func main() {
	placedAt := time.Date(2026, time.January, 10, 9, 30, 0, 0, time.UTC)
	now := time.Date(2026, time.January, 15, 9, 30, 0, 0, time.UTC)

	// 下面兩行 _ = ... 只是讓半成品能編譯；用到之後就把對應那行刪掉。
	_ = placedAt
	_ = now

	// 在這裡開始寫 ↓（TODO 1、TODO 2）
	fmt.Println("TODO: 還沒開始，把 TODO 1~2 完成後刪掉這行")
}
