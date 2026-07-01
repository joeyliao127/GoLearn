//go:build ignore

// homework-02-依賴注入｜搭配 sample-02-依賴注入
// 目標（≤5 分鐘）：定義一個 ProductRepository 介面 + 一個「假實作」，體會介面在使用方定義、實作可替換。
// 電商情境：某功能要依商品 ID 查商品名稱。我們先用一個「假的（記憶體）」實作頂著，
// 之後才換成真的資料庫——只要它符合介面，換上去就能用。
//
// 已給好：Product struct、ProductRepository 介面。你要補的是「假實作」。
// TODO 1：定義 FakeProductRepo struct（欄位 items map[int]string，代表 id → 商品名）。
// TODO 2：幫 FakeProductRepo 寫方法 GetName(id int) (string, bool)，從 items 查、回 (名稱, 是否存在)。
// 提示：Go 的 map 查詢 name, ok := r.items[id] 剛好回兩個值，直接 return 即可。
// TODO 3：在 main 建 FakeProductRepo（items 塞 1:"鍵盤"、2:"滑鼠"），指派給 var repo ProductRepository，查 id=1 印「找到：鍵盤」、查 id=99 印「查無此商品」。
// 完成後跑：go run 007-介面/homework-02-依賴注入.go
// 驗收：印出「找到：鍵盤」與「查無此商品」兩行
package main

import "fmt"

type Product struct {
	ID   int
	Name string
}

// 介面在「使用方」定義：某功能需要「依 ID 查商品名」，就宣告這個最小需求。
type ProductRepository interface {
	GetName(id int) (string, bool)
}

// TODO 1 / 2：在這裡定義 FakeProductRepo 與它的 GetName 方法

func main() {
	// 在這裡開始寫 ↓（TODO 3：建 FakeProductRepo、指派給 var repo ProductRepository、查兩個 id）
	fmt.Println("TODO: 還沒開始，把 TODO 1~3 完成後刪掉這行")
}
