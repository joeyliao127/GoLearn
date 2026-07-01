//go:build ignore

// homework-02-variadic｜搭配 sample-02-variadic
// 目標（≤5 分鐘）：寫一個帶「可變參數」的函式，幫商品組出「名稱＋標籤」字串。
// 電商情境：商品有一個名稱，加上 0 到多個標籤（熱銷、限量、免運…），標籤數量不固定。
//
// TODO 1：宣告函式 withTags(name string, tags ...string) string
// （name 是固定參數，tags 是可變參數、放最後；函式內 tags 就是一個 []string）。
// TODO 2：用 strings.Join(tags, ", ") 把標籤串起來，回傳「name（N 個標籤）：t1, t2, ...」，
// N 用 len(tags)；沒有標籤時 N=0（不用特別處理也算過關）。
// TODO 3：在 main 呼叫 withTags("機械鍵盤", "熱銷", "限量", "免運") 並用 fmt.Println 印出來。
//
// 完成後跑：go run 003-函式/homework-02-variadic.go
// 驗收：印出「機械鍵盤（3 個標籤）：熱銷, 限量, 免運」；只傳 withTags("滑鼠墊") → 「滑鼠墊（0 個標籤）：」。
package main

import (
	"fmt"
	"strings"
)

// TODO 1+2：在這裡宣告 func withTags(name string, tags ...string) string
// （提示：strings.Join([]string{"a","b"}, ", ") 會得到 "a, b"）

func main() {
	// 先擺一行示範用的呼叫；寫好 withTags 後把下面兩行換成真正的呼叫＋列印。
	_ = strings.Join // 讓 import 先不報「未使用」；寫好 withTags 用到它後就刪這行

	// 在這裡開始寫 ↓（TODO 3）
	fmt.Println("TODO: 還沒開始，把 TODO 1~3 完成後刪掉這行")
}
