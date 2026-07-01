//go:build ignore

// homework-01-傳遞｜搭配 sample-01-傳遞
// 目標（≤5 分鐘）：練習「ctx 當第一參數層層傳」——寫一個吃 ctx 的函式，並從 main 呼叫。
// 電商情境：查詢庫存。handler 收到請求，呼叫 getStock(ctx, sku) 去（假）查庫存回傳數量。
//
// TODO 1：宣告函式 getStock(ctx context.Context, sku string) int
//
//	—— ctx 一定放「第一個參數」、型別 context.Context、命名叫 ctx。
//	函式內先印一行「查詢 sku=<sku> 的庫存」，再回傳 42（假資料）。
//	（ctx 這題還用不到，函式體內加一行 _ = ctx 表示「收到但暫時沒用」即可。）
//
// TODO 2：在 main 用 context.Background() 建根 ctx，呼叫 getStock(ctx, "SKU-001")，
//
//	把回傳值用 fmt.Printf 印成「SKU-001 庫存：42 件」。
//
// 完成後跑：go run 011-context/homework-01-傳遞.go
// 驗收：印出「查詢 sku=SKU-001 的庫存」與「SKU-001 庫存：42 件」兩行
package main

import (
	"context"
	"fmt"
)

// TODO 1：在這裡宣告 func getStock(ctx context.Context, sku string) int

func main() {
	// 提示：先用 context.Background() 取得根 ctx
	ctx := context.Background()

	// 下面這行 _ = ... 只是讓半成品能編譯；寫好 getStock 並呼叫它之後就把這行刪掉。
	_ = ctx

	// 在這裡開始寫 ↓（TODO 2：呼叫 getStock 並印出結果）
	fmt.Println("TODO: 還沒開始，把 TODO 1~2 完成後刪掉這行")
}
