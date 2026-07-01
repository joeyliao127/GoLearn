//go:build ignore

// sample-03-new與make｜對應「教學.html › 8~11 節」
// 電商情境：建立各種資料——一筆商品、一個購物車（slice）、一份庫存表（map）。
// 這裡把 Go 兩個「配置記憶體」的內建函式講清楚：
//
//	new(T)      → 配置一塊「零值」記憶體，回傳 *T（指標）。任何型別都能用，但實務少用。
//	&T{...}     → struct 慣用寫法：一次給初值，回傳 *T，比 new 好用。
//	make(T,...) → 只給 slice / map / channel 用，回傳「初始化好的本體」（不是指標）。
//
// 重點：new 回指標、make 回本體；struct 用 &T{} 不用 new；map 一定要先 make/字面值才能寫。
// 跑法：go run 006-指標/sample-03-new與make.go
package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	// ── 8. new(T)：配置零值記憶體，回傳指標 *T ───────────────
	// p1 是 *Product，指向一個「零值」的 Product{Name:"", Price:0}。
	p1 := new(Product)
	p1.Name = "4K 螢幕" // p1.Name 是 (*p1).Name 的縮寫（自動解參考）
	p1.Price = 8990
	fmt.Printf("new(Product) → 型別 %T，內容 %+v\n", p1, *p1)

	// ── 9. struct 慣例：直接用 &T{...}，可同時給初值 ─────────
	// 效果和 new 一樣拿到 *Product，但能一行帶初值，所以 Go 程式幾乎都用這個、很少寫 new。
	p2 := &Product{Name: "喇叭", Price: 1200}
	fmt.Printf("&Product{...} → 型別 %T，內容 %+v\n", p2, *p2)

	// ── 10. make：只給 slice / map / channel，回傳「本體」不是指標 ──
	// slice：make([]T, len, cap)。這裡建一個空購物車，長度 0、先預留容量 4。
	cart := make([]Product, 0, 4)
	fmt.Printf("make([]Product,0,4) → 型別 %T，len=%d cap=%d\n", cart, len(cart), cap(cart))
	cart = append(cart, *p1, *p2) // 放兩件進去（append 會用到預留的容量）
	fmt.Printf("append 兩件後 → len=%d cap=%d\n", len(cart), cap(cart))

	// map：庫存表 商品名 → 數量。map 一定要先 make（或用字面值）才能寫入。
	stock := make(map[string]int)
	stock["4K 螢幕"] = 10
	stock["喇叭"] = 5
	fmt.Printf("make(map) → 4K 螢幕庫存=%d，喇叭庫存=%d\n", stock["4K 螢幕"], stock["喇叭"])

	// ── 11. 陷阱：nil map「讀」可以、「寫」會 panic（呼應主題 004）──
	// 只宣告沒 make 的 map 是 nil。讀 nil map 不會炸（回零值），但寫入會 panic。
	var priceTable map[string]int // nil map（沒 make）
	fmt.Printf("nil map 讀取 OK → 得到零值 %d（不會 panic）\n", priceTable["不存在"])
	// priceTable["鍵盤"] = 2490            // ← 若解開這行會 panic: assignment to entry in nil map
	priceTable = make(map[string]int) // 先 make 才能安全寫入
	priceTable["鍵盤"] = 2490
	fmt.Printf("make 之後才寫入 → 鍵盤=%d\n", priceTable["鍵盤"])

	// 小結：要不要指標看「會不會被大量修改 / 要不要共享同一份」；
	// slice/map/channel 用 make，其餘 struct 用 &T{}，new 幾乎用不到。
	fmt.Println("小結：struct 用 &T{}，slice/map/chan 用 make，new 幾乎用不到")
}
