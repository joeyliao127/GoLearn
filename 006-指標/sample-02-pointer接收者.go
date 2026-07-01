//go:build ignore

// sample-02-pointer接收者｜對應「教學.html › 5~7 節」
// 電商情境：訂單的方法——「打折」要改到訂單本身 → 用 pointer receiver (*Order)；
// 「算總額」只是讀、不改東西 → 用 value receiver (Order) 就好。
// 這是主題 005 學過的「方法/receiver」的續集：現在補上「receiver 該用值還是指標」。
// 重點：① 要改到自身的方法用 *T；② 全值或全指標別混用；③ range 拿到的是副本、要改原值得用索引。
// 跑法：go run 006-指標/sample-02-pointer接收者.go
package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

type Order struct {
	ID    int
	Items []Product
}

// ── 5. pointer receiver：要「改到自己」的方法都用 *Order ──────
// receiver 寫成 (o *Order)，方法內對 o 的修改會寫回本尊。
// 注意迴圈用「索引」o.Items[i] 去改；若寫 for _, item := range，item 是副本，改了沒用（見第 7 節）。
func (o *Order) ApplyDiscount(percent int) {
	for i := range o.Items {
		o.Items[i].Price = o.Items[i].Price * (100 - percent) / 100
	}
}

// AddItem 也是會改到自身（多一件商品）→ 一樣用 pointer receiver。
func (o *Order) AddItem(p Product) {
	o.Items = append(o.Items, p)
}

// ── 6. value receiver：只「讀」不「改」→ 用 Order 就好 ────────
// Total 只是加總，不改任何欄位，用 value receiver 更安全（不怕誤改）也語意清楚。
func (o Order) Total() int {
	sum := 0
	for _, item := range o.Items {
		sum += item.Price
	}
	return sum
}

func main() {
	// 用 &Order{} 建立，拿到的是 *Order；對它呼叫方法最直接。
	order := &Order{ID: 1}
	order.AddItem(Product{Name: "機械鍵盤", Price: 2000})
	order.AddItem(Product{Name: "無線滑鼠", Price: 1000})

	// ── 呼叫 value receiver 方法：只是讀 ─────────────────────
	fmt.Printf("打折前 Total = %d\n", order.Total())

	// ── 呼叫 pointer receiver 方法：改到本尊 ─────────────────
	order.ApplyDiscount(10) // 打 9 折，改的是 order 指向的那筆訂單
	fmt.Printf("打 9 折後 Total = %d\n", order.Total())
	for _, item := range order.Items {
		fmt.Printf("  - %s：%d 元\n", item.Name, item.Price)
	}

	// ── 補充：即使 order 是「值」而非指標，也能呼叫 *Order 方法 ──
	// Go 的貼心：對「可定址(addressable)」的變數呼叫指標方法時，會自動幫你取地址
	// （val.ApplyDiscount() 自動變成 (&val).ApplyDiscount()）。所以下面兩種寫法都行。
	val := Order{ID: 2, Items: []Product{{Name: "螢幕", Price: 5000}}}
	val.ApplyDiscount(20) // 自動變 (&val).ApplyDiscount(20)，改得到 val
	fmt.Printf("值變數 val 打 8 折後 Total = %d\n", val.Total())

	// ── 7. 陷阱回顧：range 的元素是「副本」，直接改沒用 ──────
	// 這裡故意用 for _, item := range 改價格 → 只改到副本 item，原訂單不動。
	demo := &Order{ID: 3, Items: []Product{{Name: "耳機", Price: 800}}}
	for _, item := range demo.Items {
		item.Price = 0 // ← 改的是副本，白改
	}
	fmt.Printf("（陷阱）用 range 副本改價後，Total 仍是 %d（沒被改到）\n", demo.Total())
}
