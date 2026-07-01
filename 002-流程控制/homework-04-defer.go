//go:build ignore

// homework-04-defer｜搭配 sample-04-defer
// 目標（≤5 分鐘）：用 defer 確保「訂單處理結束」這行一定最後才印出來。
// 電商情境：processOrder 處理過程中會印好幾行，但「訂單處理結束」要保證壓在最後
// （就算中途有 return 也一樣）——這正是 defer 的用途。
//
// TODO 1：在 processOrder 函式「一開始」用 defer 登記印出「訂單處理結束」
// （提示：defer fmt.Println("訂單處理結束") 會等函式 return 時才執行）
// TODO 2：（觀察用，不用改）跑跑看，確認「訂單處理結束」出現在「扣款完成」之後，即使它寫在函式最前面。
//
// 完成後跑：go run 002-流程控制/homework-04-defer.go
// 驗收：輸出依序為「開始處理訂單 2001」「扣款完成」「訂單處理結束」（最後一行靠 defer）
package main

import "fmt"

func processOrder(orderID int) {
	// TODO 1：在這裡用 defer 登記印出「訂單處理結束」

	fmt.Printf("開始處理訂單 %d\n", orderID)
	fmt.Println("扣款完成")
	// 函式在這裡 return，defer 才會執行 → 印出「訂單處理結束」
}

func main() {
	processOrder(2001)
	fmt.Println("TODO: 完成 processOrder 裡的 defer 後，把這行刪掉")
}
