//go:build ignore

// homework-04-embedding｜搭配 sample-04-embedding
// 目標（≤5 分鐘）：用「嵌入(embedding)」讓 Product 共用 Audit 的稽核欄位。
// 電商情境：商品也要記錄「誰、何時建立」，直接把現成的 Audit 嵌進 Product 共用。
//
// TODO 1：在 Product 裡「嵌入」Audit（只寫型別名 Audit、不給欄位名，放在欄位列表最上面）
// TODO 2：在 main 建立 Product 時，用 Audit: Audit{CreatedBy:"admin", CreatedAt:"2026-07-01"} 填稽核資料，
//
//	其餘欄位：Name="無線滑鼠"、Price=590
//
// TODO 3：直接用「晉升欄位」印出建立者：fmt.Println(p.CreatedBy)（不必寫 p.Audit.CreatedBy）
//
// 完成後跑：go run 005-struct與方法/homework-04-embedding.go
// 驗收：印出「建立者：admin」（靠嵌入後的欄位晉升，直接 p.CreatedBy 取得）
package main

import "fmt"

type Audit struct {
	CreatedBy string
	CreatedAt string
}

// TODO 1：在 Product 裡嵌入 Audit（加一行只有型別名的 Audit）
type Product struct {
	Name  string
	Price int
}

func main() {
	// 在這裡開始寫 ↓（TODO 2：建立 Product 並填 Audit；TODO 3：印 p.CreatedBy）
	fmt.Println("TODO: 還沒開始，把 TODO 1~3 完成後刪掉這行")
}
