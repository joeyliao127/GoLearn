//go:build ignore

// sample-02-時間｜對應「教學.html › 5~8 節」
// 電商情境：訂單有「下單時間」，我們常要：格式化成字串存進 DB／回給前端（RFC3339）、
// 算「下單到現在幾天」、算「7 天後到期日」、以及與前端交換用的「Unix 時間戳」。
// 重點：① time.Time 是時間的型別；② 用「參考時間」版面格式化（不是 yyyy-MM-dd）；
// ③ 加減用 AddDate / Add；④ 兩時間相減 Sub 得 Duration；⑤ Unix 時間戳互轉。
// 註：本檔刻意用「固定的基準時間」而非 time.Now()，讓每次輸出都一樣、方便對照驗收。
// 跑法：go run 012-JSON與時間/sample-02-時間.go
package main

import (
	"fmt"
	"time"
)

func main() {
	// ── 5 節：建立一個 time.Time（用固定基準時間，輸出才可預期）──────
	// time.Date(年,月,日,時,分,秒,奈秒,時區)。用 time.UTC 固定時區，避免受本機時區影響。
	placedAt := time.Date(2026, time.January, 10, 9, 30, 0, 0, time.UTC)
	fmt.Println("下單時間（預設印法）：", placedAt)

	// ── 6 節：格式化 —— Go 的「參考時間」版面 ──────────────────────
	// Go 不用 yyyy/MM/dd，而是用「一個固定的參考時間」當範本：
	//   2006-01-02 15:04:05（記法：1 2 3 4 5 6＝月日時分秒年，外加時區 -0700）。
	// 你想要什麼格式，就把這個參考時間「排成那個樣子」當 layout 傳給 Format。
	fmt.Println("自訂格式：", placedAt.Format("2006/01/02 15:04"))

	// RFC3339 是 API／DB 最常用的標準格式（time 套件有內建常數，不用自己記版面）。
	fmt.Println("RFC3339：", placedAt.Format(time.RFC3339))

	// 反向：把 RFC3339 字串 Parse 回 time.Time（回傳 (time.Time, error)）。
	parsed, err := time.Parse(time.RFC3339, "2026-01-17T09:30:00Z")
	if err != nil {
		fmt.Println("parse 失敗：", err)
		return
	}

	// ── 7 節：加減時間 ────────────────────────────────────────────
	// AddDate(年,月,日)：算「7 天後到期」最直覺。
	expireAt := placedAt.AddDate(0, 0, 7)
	fmt.Println("下單 +7 天到期：", expireAt.Format(time.RFC3339))

	// Add 吃 Duration（時分秒等級）：例如「保留庫存 30 分鐘」。
	holdUntil := placedAt.Add(30 * time.Minute)
	fmt.Println("庫存保留到：", holdUntil.Format(time.RFC3339))

	// 比較先後：Before / After / Equal（別用 == 比 time.Time，時區/單調時鐘會有坑）。
	fmt.Println("到期是否晚於下單？", expireAt.After(placedAt)) // true

	// ── 8 節：兩時間相減 Sub → Duration ；以及 Unix 時間戳 ─────────
	// parsed 是 1/17、placedAt 是 1/10，相差 7 天。Sub 得到 Duration。
	elapsed := parsed.Sub(placedAt)
	fmt.Printf("下單到收貨經過：%v（= %.0f 小時 = %.0f 天）\n",
		elapsed, elapsed.Hours(), elapsed.Hours()/24)

	// Unix 時間戳：從 1970-01-01 UTC 起算的「秒數」，是跨語言/前後端交換時間的通用格式。
	ts := placedAt.Unix()
	fmt.Println("下單的 Unix 時間戳（秒）：", ts)

	// 從時間戳還原成 time.Time：time.Unix(秒, 奈秒)。這裡再轉 UTC 讓印出來穩定。
	restored := time.Unix(ts, 0).UTC()
	fmt.Println("由時間戳還原：", restored.Format(time.RFC3339))
}
