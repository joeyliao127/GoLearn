//go:build ignore

// sample-02-方法｜對應「教學.html › 4~6 節」
// 電商情境：幫 Order / Product 掛上「方法」，把「算訂單總額」「產生商品標籤」這種
// 邏輯，綁在資料旁邊——這正是電商 repo 裡 model 帶行為的樣子。
// 重點：① 方法 = 帶「receiver」的函式；② 這裡全用 value receiver（(o Order)），
// 適合「只讀不改」的計算；③ receiver 取什麼名字都行，慣例用型別首字母小寫。
// 跑法：go run 005-struct與方法/sample-02-方法.go
package main

import "fmt"

type Product struct {
	ID    string
	Name  string
	Price int
}

type Order struct {
	ID      string
	Product Product
	Qty     int
}

// ── 4. 方法：在 func 和函式名之間，多一個 (receiver) ─────
// 讀作「Order 型別有一個 Total 方法」。receiver o 就像其他語言的 this / self，
// 但 Go 不用 this 這種關鍵字，而是自己命名（慣例：型別首字母小寫，這裡 o）。
// value receiver：呼叫時 o 是「這個 Order 的一份複製」，適合純計算、不改本體。
func (o Order) Total() int {
	return o.Product.Price * o.Qty
}

// ── 5. 方法就是「綁在型別上的函式」，可有回傳值、可讀 receiver 的欄位 ──
// Label 用 Product 自己的欄位組一段顯示字串（例如商品列表的標籤）。
func (p Product) Label() string {
	return fmt.Sprintf("%s（NT$%d）", p.Name, p.Price)
}

// 方法也可以帶「額外參數」，跟一般函式一樣。
// discounted 回傳「打折後的總額」，但注意：它只是算出新數字回傳，
// 沒有去改 o 本身（value receiver 改 o 也只是改複製品，白改）。
func (o Order) TotalWithDiscount(percent int) int {
	total := o.Total()
	return total * (100 - percent) / 100
}

// ── 6. 陷阱示範：value receiver 改欄位「改不到本體」 ──────
// 這個方法想把數量翻倍，但 receiver 是複製品，改完 return 就丟了，外面的 Order 沒變。
// （想「改得到本體」要用 pointer receiver：(o *Order)——那是主題 006 的主角，這裡先知道有這回事。）
func (o Order) doubleQtyWrong() {
	o.Qty *= 2 // 只改到複製品 o，函式一結束就消失
}

func main() {
	order := Order{
		ID:      "O-20260701-002",
		Product: Product{ID: "P001", Name: "機械鍵盤", Price: 2490},
		Qty:     3,
	}

	// 用「值.方法()」呼叫，跟其他語言的物件方法呼叫長得一樣。
	fmt.Printf("商品標籤：%s\n", order.Product.Label())
	fmt.Printf("訂單總額：%d 元\n", order.Total())
	fmt.Printf("打 9 折（percent=10）後：%d 元\n", order.TotalWithDiscount(10))

	// 印證第 6 節的陷阱：呼叫 doubleQtyWrong 後，Qty 仍是 3（沒被改到）。
	order.doubleQtyWrong()
	fmt.Printf("呼叫 doubleQtyWrong 後 Qty 仍是 %d（value receiver 改不到本體）\n", order.Qty)
}
