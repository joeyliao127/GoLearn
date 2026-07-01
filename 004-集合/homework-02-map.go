//go:build ignore

// homework-02-map｜搭配 sample-02-map
// 目標（≤5 分鐘）：用 map[string]int 統計「每個分類各有幾個商品」。
// 電商情境（商品清單已給好）：後台要顯示各分類的商品數量。
//
// TODO 1：用 make 建一個 counts := make(map[string]int)（分類名 → 數量）
// TODO 2：用 for range 走訪 products，每一項就 counts[p.Category]++
// （關鍵：對「還不存在的分類」直接 ++ 是安全的——讀到零值 0，加 1 後變成 1）
// TODO 3：走訪 counts，用 fmt.Printf 印出「分類 X：Y 個」
//
// 完成後跑：go run 004-集合/homework-02-map.go
// 驗收：3C=3 個、周邊=2 個、耗材=1 個（順序不保證，數字對就好）
package main

import "fmt"

type Product struct {
	Name     string
	Category string
}

func main() {
	products := []Product{
		{Name: "機械鍵盤", Category: "3C"},
		{Name: "無線滑鼠", Category: "3C"},
		{Name: "4K 螢幕", Category: "3C"},
		{Name: "滑鼠墊", Category: "周邊"},
		{Name: "手機支架", Category: "周邊"},
		{Name: "清潔布", Category: "耗材"},
	}

	// 下面這行 _ = ... 只是讓半成品能編譯；開始用 products 之後就把這行刪掉。
	_ = products

	// 在這裡開始寫 ↓
	fmt.Println("TODO: 還沒開始，把 TODO 1~3 完成後刪掉這行")
}
