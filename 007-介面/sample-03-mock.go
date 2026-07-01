//go:build ignore

// sample-03-mock｜對應「教學.html › 8~10 節」
// 電商情境：想驗證「下單服務」的邏輯對不對，但又不想真的連資料庫。做法：注入一個
// 「mock（假的）repository」——它不存真資料，只負責①記錄被呼叫了幾次、②可被指定「故意失敗」，
// 讓我們專心測 service 自己的邏輯。這正是主題 013 測試會用到的手法（這裡先手動示範）。
// 重點：① mock 也是靠「實作介面」；② 用 mock 觀察互動 / 注入錯誤；③ nil interface 陷阱。
// 跑法：go run 007-介面/sample-03-mock.go
package main

import (
	"errors"
	"fmt"
)

// 沿用 sample-02 的角色：一筆訂單 + service 依賴的 repository 介面。
type Order struct {
	ID     int
	Amount int
}

type OrderRepository interface {
	Save(o Order) error
}

// 被測對象：下單服務。金額 <= 0 擋下來，否則呼叫 repo.Save。
type OrderService struct {
	repo OrderRepository
}

func (s *OrderService) PlaceOrder(o Order) error {
	if o.Amount <= 0 {
		return fmt.Errorf("金額必須大於 0")
	}
	return s.repo.Save(o)
}

// ── 8. Mock：一個「為測試而生」的假實作 ─────────────────────
// 它同樣實作 OrderRepository（有 Save），但不存真資料，而是：
//   - SaveCalls：記錄 Save 被呼叫幾次（驗證「有沒有真的去存」）
//   - LastSaved：記下最後存的內容（驗證「存進去的資料對不對」）
//   - FailWith ：若不為 nil，Save 就回這個錯（模擬資料庫爆掉，測 service 的錯誤處理）
type MockOrderRepo struct {
	SaveCalls int
	LastSaved Order
	FailWith  error
}

func (m *MockOrderRepo) Save(o Order) error {
	m.SaveCalls++ // 記錄互動
	m.LastSaved = o
	if m.FailWith != nil {
		return m.FailWith // 被指定要失敗 → 回注入的錯
	}
	return nil
}

func main() {
	// ── 9-a. 用 mock 驗證「正常下單」：service 有沒有正確呼叫 Save ──
	mock := &MockOrderRepo{}
	svc := &OrderService{repo: mock} // 注入 mock，而不是真 DB

	_ = svc.PlaceOrder(Order{ID: 1001, Amount: 1500})

	// 「斷言」：手動檢查 mock 記到的互動符不符合預期（013 會用 testing 套件自動化）。
	fmt.Printf("Save 被呼叫 %d 次（預期 1）\n", mock.SaveCalls)
	fmt.Printf("最後存入：#%d %d 元（預期 #1001 1500 元）\n", mock.LastSaved.ID, mock.LastSaved.Amount)

	// ── 9-b. 用 mock「注入錯誤」：測 service 對 DB 失敗的反應 ──
	failing := &MockOrderRepo{FailWith: errors.New("DB 連線逾時")}
	svc2 := &OrderService{repo: failing}
	if err := svc2.PlaceOrder(Order{ID: 1002, Amount: 999}); err != nil {
		fmt.Println("如預期收到錯誤：", err)
	}

	// ── 9-c. 金額不合法：service 應在呼叫 Save 之前就擋下 ──
	guard := &MockOrderRepo{}
	svc3 := &OrderService{repo: guard}
	_ = svc3.PlaceOrder(Order{ID: 1003, Amount: 0})
	fmt.Printf("金額為 0 時 Save 被呼叫 %d 次（預期 0，代表被擋下）\n", guard.SaveCalls)

	// ── 10. 陷阱：nil interface ≠ nil ────────────────────────
	// 介面值內部其實有兩格：(型別, 值)。只有「兩格都空」時介面才等於 nil。
	// 把一個「型別已知、但值為 nil 的指標」塞進介面，介面就「非 nil」了——超容易誤判。
	var repo OrderRepository // 什麼都沒放：型別和值都空 → 真正的 nil
	fmt.Println("--- nil interface 陷阱 ---")
	fmt.Printf("空介面 == nil ? %v（預期 true）\n", repo == nil)

	var p *MockOrderRepo = nil // 一個「值為 nil」的具體指標
	repo = p                   // 塞進介面：型別=*MockOrderRepo（已知）、值=nil
	// 現在介面「帶著型別資訊」，所以它不等於 nil，即使底下的指標是 nil！
	fmt.Printf("裝了 nil 指標的介面 == nil ? %v（預期 false，這就是陷阱）\n", repo == nil)
}
