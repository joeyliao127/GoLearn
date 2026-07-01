//go:build ignore

// sample-01-傳遞｜對應「教學.html › 1~3 節」
// 電商情境：一個「查詢訂單」的請求進來，要一路呼叫 handler → service → repository。
// 這條呼叫鏈上的每一層，都把同一個 ctx 當「第一個參數」往下傳。
// 重點：
//  1. 為什麼要 context（給整條呼叫鏈一個共同的「取消 / 逾時 / 帶值」訊號）。
//  2. ctx 放第一參數、命名固定叫 ctx、型別是 context.Context。
//  3. 最上層（main / 收到請求處）用 context.Background() 當「根 context」。
//
// 跑法：go run 011-context/sample-01-傳遞.go
package main

import (
	"context"
	"fmt"
)

// 一筆訂單（金額用 int 存「元」，別用 float）。
type Order struct {
	ID     int
	Amount int
}

func main() {
	// ── 1. 為何要 context：先看「根 context」怎麼來 ────────────
	// context 是一個沿呼叫鏈往下傳的物件，攜帶三種東西：取消訊號、逾時、request 範圍的值。
	// 最上層要先有一個「根」。慣例：main、程式初始化、測試、以及「剛收到一個請求」的地方，
	// 都用 context.Background()——它永不取消、沒有值、沒有逾時，是一切 ctx 的起點。
	ctx := context.Background()

	// 模擬一個進來的請求，交給 handler 處理（真的 web 框架會幫你把 ctx 準備好）。
	handleGetOrder(ctx, 1001)
}

// ── 2. ctx 當「第一個參數」，型別 context.Context，命名固定叫 ctx ──
// 這是 Go 鐵一般的慣例：只要一個函式「可能要被取消 / 逾時 / 讀 request 範圍值」，
// 就把 ctx 放在參數列「最前面」。web handler 幾乎必吃 ctx，是最典型的例子。
func handleGetOrder(ctx context.Context, orderID int) {
	fmt.Printf("[handler] 收到查詢訂單 %d 的請求\n", orderID)

	// 往下呼叫 service 時，把「同一個 ctx」原封不動傳下去（先不加工，第 2、3 個 sample 才會加工）。
	order, err := getOrderService(ctx, orderID)
	if err != nil {
		fmt.Println("[handler] 查詢失敗：", err)
		return
	}
	fmt.Printf("[handler] 回應：訂單 %d，金額 %d 元\n", order.ID, order.Amount)
}

// ── 3. 中間層：service 收到 ctx，繼續往下傳給 repository ────────
// 這一層自己不一定用到 ctx，但它「有義務」把 ctx 繼續傳下去，
// 讓最底層（DB 查詢）也能收到取消 / 逾時訊號。這就是「層層傳遞」的意義。
func getOrderService(ctx context.Context, orderID int) (Order, error) {
	fmt.Println("  [service] 處理商業邏輯，把 ctx 傳給 repository")
	return getOrderFromDB(ctx, orderID)
}

// 最底層：模擬「查資料庫」。真實世界這裡會做 DB 查詢，而 database/sql 的
// QueryContext 之類 API 全都吃 ctx——所以 ctx 一定要能傳到這麼深。
func getOrderFromDB(ctx context.Context, orderID int) (Order, error) {
	// 這裡先示範「ctx 有到手」；真正用 ctx 做取消 / 逾時是下一個 sample 的主題。
	// _ = ctx 表示「我收到了 ctx，暫時還沒用它做事」，避免未使用參數的誤會。
	_ = ctx
	fmt.Println("    [db] 收到 ctx，執行查詢（假資料）")
	return Order{ID: orderID, Amount: 1980}, nil
}
