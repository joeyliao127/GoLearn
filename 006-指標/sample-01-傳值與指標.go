//go:build ignore

// sample-01-傳值與指標｜對應「教學.html › 1~4 節」
// 電商情境：對一筆「訂單」加商品——用「傳值」改不到本尊，用「傳指標」才改得到。
// 核心心法（整個主題就靠這句）：Go 永遠是「傳值（pass by value）」——傳進函式的都是一份 copy。
// 想改到呼叫端的本尊，就把「地址」傳進去，讓兩邊指向同一塊記憶體。
// 順便看懂三個符號：*T（指標型別）、&x（取地址）、*p（解參考，讀/寫本尊）。
// 跑法：go run 006-指標/sample-01-傳值與指標.go
package main

import "fmt"

// Product 商品：一筆相對「唯讀、不太會變」的小資料。
type Product struct {
	Name  string
	Price int // 用整數存「元」，別用 float 存錢
}

// Order 訂單：會「一直被修改」的資料（加商品、改狀態）——這種最需要指標。
type Order struct {
	ID     int
	Items  []Product
	Status string
}

func main() {
	keyboard := Product{Name: "機械鍵盤", Price: 2490}
	mouse := Product{Name: "無線滑鼠", Price: 590}

	// ── 1. 傳值心法：函式收到的是「一份 copy」，改不到本尊 ─────
	// addItemByValue 收到的 o 只是 orderA 的複製品，函式裡動的全是 copy，
	// 函式 return 後 copy 就被丟掉，orderA 完全沒變。
	orderA := Order{ID: 1, Status: "new"}
	addItemByValue(orderA, keyboard)
	fmt.Printf("[傳值]   加了商品後 → 件數=%d 狀態=%q（沒變！改的是 copy）\n",
		len(orderA.Items), orderA.Status)

	// ── 2. 傳指標：用 & 取地址，函式透過地址寫回本尊 ──────────
	// &orderB 是「指向 orderB 的地址」，型別是 *Order。
	// 函式透過這個地址寫入，動到的就是 orderB 本尊。
	orderB := Order{ID: 2, Status: "new"}
	addItemByPointer(&orderB, keyboard) // & = 取地址
	addItemByPointer(&orderB, mouse)
	fmt.Printf("[傳指標] 加了商品後 → 件數=%d 狀態=%q（真的變了）\n",
		len(orderB.Items), orderB.Status)

	// ── 3. 三符號速記：*T 型別、&x 取地址、*p 解參考 ──────────
	price := 100
	var p *int = &price // p 的型別是 *int（指向 int 的指標）；&price 取 price 的地址
	fmt.Printf("price=%d，p 指向的地址=%p，*p 讀回本尊=%d\n", price, p, *p)
	*p = 250 // *p = ... 透過指標「寫」本尊
	fmt.Printf("經過 *p = 250 後，price=%d（本尊被改了）\n", price)

	// ── 4. 「傳指標」其實也是傳值：copy 的是「地址」 ──────────
	// 別誤會成「指標就不 copy 了」。指標本身也是一個值（一串地址數字），
	// 傳進函式一樣被複製一份；只是這份 copy 和原本的指標「指向同一塊記憶體」，
	// 所以透過它寫入才會動到本尊。下面印出兩個地址相同可證明。
	fmt.Printf("orderB 的地址：外面 &orderB=%p\n", &orderB)
	showAddr(&orderB) // 函式裡收到的地址值和外面一模一樣
}

// ── 反例：用「傳值」想改訂單 → 改不到 ──────────────────────
// o 是 orderA 的一份 copy；append、改 Status 都只動到 copy。
func addItemByValue(o Order, p Product) {
	o.Items = append(o.Items, p)
	o.Status = "updated(其實在改 copy)"
}

// ── 正解：用「傳指標」改訂單 → 改得到 ──────────────────────
// o 是 *Order（指向本尊的地址）。o.Items 是 (*o).Items 的縮寫——
// 對指標取欄位時，Go 會自動幫你解參考(dereference)，不用寫成 (*o).Items。
func addItemByPointer(o *Order, p Product) {
	o.Items = append(o.Items, p)
	o.Status = "updated"
}

// showAddr 印出「函式內收到的地址」，用來證明傳指標時 copy 的是地址、指向同一塊。
func showAddr(o *Order) {
	fmt.Printf("orderB 的地址：函式內 o=%p（和外面一樣 → 指向同一筆訂單）\n", o)
}
