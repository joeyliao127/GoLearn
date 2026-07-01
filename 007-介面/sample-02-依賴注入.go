//go:build ignore

// sample-02-依賴注入｜對應「教學.html › 5~7 節」
// 電商情境：OrderService（下單服務）需要「存取訂單資料」，但它不該在乎資料是存在
// 記憶體、MySQL 還是遠端 API。做法：讓 service 依賴一個 OrderRepository「介面」，
// 真正的實作由外部「注入」進來——這就是依賴注入（DI）。
// 重點：① 介面在「使用方（service）」定義；② service 只認介面、不 new 具體實作；
// ③ 換實作（記憶體→DB）時，service 一行都不用改。
// 跑法：go run 007-介面/sample-02-依賴注入.go
package main

import "fmt"

// 一筆訂單（金額用 int 存「元」，別用 float）。
type Order struct {
	ID     int
	Amount int
}

// ── 5. 介面「在使用方定義」：service 需要什麼，就宣告什麼 ─────
// OrderRepository 是 OrderService 對「資料層」的最小需求：存一筆、依 ID 找一筆。
// Go 慣例：介面定義在「用它的人」這邊（consumer），而不是實作那邊。
type OrderRepository interface {
	Save(o Order) error
	FindByID(id int) (Order, bool)
}

// ── 6. service 只依賴介面，不依賴任何具體型別 ─────────────────
// 欄位型別是介面 OrderRepository；OrderService 完全不知道背後是誰在存資料。
type OrderService struct {
	repo OrderRepository
}

// 慣例：用一個建構函式把依賴「注入」進來（constructor injection）。
// 參數收介面 → 呼叫端要塞什麼實作都行（記憶體版、DB 版、測試用的 mock…）。
func NewOrderService(repo OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

// PlaceOrder 是業務邏輯：它只透過介面方法做事，不碰底層細節。
func (s *OrderService) PlaceOrder(o Order) error {
	if o.Amount <= 0 {
		return fmt.Errorf("訂單金額必須大於 0，收到 %d", o.Amount)
	}
	return s.repo.Save(o)
}

// ── 7. 一個具體實作：記憶體版 repository（用 map 當資料庫）────────
// 它「剛好」有 Save / FindByID 兩個方法 → 隱式滿足 OrderRepository。
type InMemoryOrderRepo struct {
	data map[int]Order
}

func NewInMemoryOrderRepo() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{data: make(map[int]Order)}
}

func (r *InMemoryOrderRepo) Save(o Order) error {
	r.data[o.ID] = o
	fmt.Printf("[記憶體] 已存訂單 #%d（%d 元）\n", o.ID, o.Amount)
	return nil
}

func (r *InMemoryOrderRepo) FindByID(id int) (Order, bool) {
	o, ok := r.data[id]
	return o, ok
}

func main() {
	// 組裝：先建一個「記憶體版 repo」，再注入給 service。
	// 之後若要換成真的資料庫，只要換這一行 new 的實作，service 完全不動。
	repo := NewInMemoryOrderRepo()
	service := NewOrderService(repo)

	// 走一次業務流程：下單 → service 透過介面存資料。
	if err := service.PlaceOrder(Order{ID: 1001, Amount: 1500}); err != nil {
		fmt.Println("下單失敗：", err)
	}

	// 讀回來確認。
	if o, ok := repo.FindByID(1001); ok {
		fmt.Printf("查到訂單 #%d：%d 元\n", o.ID, o.Amount)
	}

	// 金額不合法的情況：業務邏輯擋下來、不會存進 repo。
	if err := service.PlaceOrder(Order{ID: 1002, Amount: 0}); err != nil {
		fmt.Println("下單失敗：", err)
	}
}
