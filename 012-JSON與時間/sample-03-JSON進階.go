//go:build ignore

// sample-03-JSON進階｜對應「教學.html › 9~11 節」
// 電商情境：真實的訂單 JSON 是「巢狀」的——一筆 Order 底下有客戶物件、有一串商品明細（slice）、
// 還有一個 time.Time 下單時間。另外，別人傳來的 JSON 常有我們沒定義的欄位，
// 這時用 map[string]any 就能「照單全收」不漏資料。
// 重點：① 巢狀 struct / slice 的序列化；② time.Time 進 JSON 自動變 RFC3339 字串；
// ③ 未知結構用 map[string]any 接，取值要用「型別斷言」還原型別。
// 跑法：go run 012-JSON與時間/sample-03-JSON進階.go
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// ── 9 節：巢狀 struct + slice ─────────────────────────────────
// Order 內嵌 Customer（物件）與 []Item（陣列）；巢狀就是「struct 欄位又是 struct/slice」。
type Customer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Item struct {
	Name  string `json:"name"`
	Price int    `json:"price"` // 金額用 int
	Qty   int    `json:"qty"`
}

type Order struct {
	ID       string    `json:"id"`
	Customer Customer  `json:"customer"`  // 巢狀物件
	Items    []Item    `json:"items"`     // 陣列（會變成 JSON array）
	PlacedAt time.Time `json:"placed_at"` // time.Time → 自動序列化成 RFC3339 字串
}

func main() {
	// ── 9~10 節：組一筆巢狀訂單並 Marshal ─────────────────────────
	// 用固定時間讓輸出可預期（沿用 sample-02 的「參考時間」觀念）。
	o := Order{
		ID:       "A-3001",
		Customer: Customer{Name: "王小明", Email: "ming@example.com"},
		Items: []Item{
			{Name: "藍牙耳機", Price: 1290, Qty: 1},
			{Name: "USB-C 線", Price: 199, Qty: 2},
		},
		PlacedAt: time.Date(2026, time.January, 10, 9, 30, 0, 0, time.UTC),
	}

	pretty, _ := json.MarshalIndent(o, "", "  ")
	fmt.Println("巢狀訂單 → JSON：")
	fmt.Println(string(pretty))
	// 注意 placed_at 那行：time.Time 不用你動手，就變成 "2026-01-10T09:30:00Z"。

	// 在 Go 這端把巢狀資料算一算：訂單總金額。
	total := 0
	for _, it := range o.Items {
		total += it.Price * it.Qty
	}
	fmt.Println("訂單總金額 =", total, "元") // 1290*1 + 199*2 = 1688

	// ── 11 節：未知結構用 map[string]any 接 ───────────────────────
	// 情境：第三方 webhook 傳來一包我們「沒定義 struct」的 JSON，欄位還可能變動。
	// 這時把它 Unmarshal 進 map[string]any（any 就是 interface{}，Go 1.18+ 的別名）。
	raw := []byte(`{
		"event": "payment.captured",
		"amount": 1688,
		"paid": true,
		"tags": ["vip", "urgent"],
		"meta": {"gateway": "tappay", "retry": 0}
	}`)

	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		fmt.Println("unmarshal 失敗：", err)
		return
	}

	// 取值要用「型別斷言」把 any 還原成具體型別。
	// 陷阱：JSON 的「數字」不管整數小數，一律被解成 float64！所以 amount 要用 float64 接。
	event := m["event"].(string)
	amount := m["amount"].(float64) // 不能寫 .(int)，會 panic
	paid := m["paid"].(bool)
	fmt.Printf("事件=%s、金額=%.0f、已付款=%v\n", event, amount, paid)

	// 巢狀進去：陣列是 []any、物件是 map[string]any，要一層層斷言。
	tags := m["tags"].([]any)
	fmt.Println("第一個標籤：", tags[0].(string))

	meta := m["meta"].(map[string]any)
	fmt.Println("金流商：", meta["gateway"].(string))
}
