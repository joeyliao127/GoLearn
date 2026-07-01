//go:build ignore

// sample-03-sentinel與自訂｜對應「教學.html › 7~9 節」
// 電商情境：查詢訂單。找不到訂單時，我們想讓呼叫方不只知道「找不到」，
// 還能拿到「找不到的是哪種資源、哪個 ID」——這時固定字串的 sentinel 不夠用，
// 要自訂一個「帶欄位的 error 型別」。
// 重點：① sentinel error：var ErrXxx = errors.New(...)，給「不需帶資料」的固定錯誤；
// ② 自訂 error 型別：一個 struct 實作 Error() string，就是合法的 error，可夾帶欄位；
// ③ errors.As：把錯誤鏈裡「某個型別」挖出來、還原成具體型別，好讀它的欄位。
// 跑法：go run 008-錯誤處理/sample-03-sentinel與自訂.go
package main

import (
	"errors"
	"fmt"
)

// ── 7. sentinel error：不需帶資料的固定錯誤，用 == / errors.Is 比 ──
// 慣例：套件層級變數、名字 ErrXxx、用 errors.New 建立。付款被拒就是「有或沒有」，
// 不必夾帶欄位，很適合用 sentinel。
var ErrPaymentDeclined = errors.New("付款被拒")

// ── 8. 自訂 error 型別：struct + Error() string，就能夾帶欄位 ──
// error 是介面，只要有 Error() string 方法就算實作了它（隱式，不用寫 implements）。
// 這裡 NotFoundError 帶了「資源種類」與「ID」兩個欄位，比純字串資訊多很多。
type NotFoundError struct {
	Resource string // 例如 "訂單"
	ID       string // 例如 "ORD-9999"
}

// 實作 error 介面。慣例：訊息小寫開頭、不加句點。
func (e *NotFoundError) Error() string {
	return fmt.Sprintf("找不到%s：%s", e.Resource, e.ID)
}

// findOrder：查得到回 (訂單摘要, nil)；查不到回自訂的 *NotFoundError（用指標，見教學說明）。
func findOrder(id string) (string, error) {
	orders := map[string]string{"ORD-1001": "MacBook ×1"}
	if summary, ok := orders[id]; ok {
		return summary, nil
	}
	// 回傳「帶資料」的錯誤：呼叫方之後能挖出 Resource / ID。
	return "", &NotFoundError{Resource: "訂單", ID: id}
}

func main() {
	// 先做一次成功的查詢當對照。
	if summary, err := findOrder("ORD-1001"); err == nil {
		fmt.Println("查到訂單：", summary)
	}

	// 查一個不存在的訂單，並「包一層」模擬真實呼叫鏈（外層加上是哪個 API 的脈絡）。
	_, err := findOrder("ORD-9999")
	err = fmt.Errorf("查詢訂單 API 失敗: %w", err) // 用 %w 包住，As 仍挖得到

	// ── 9. errors.As：從錯誤鏈挖出「某個具體型別」，讀它的欄位 ──
	// Is 回答「是不是某個值」；As 回答「鏈上有沒有某個型別，有就給我那個值」。
	// 用法：宣告一個目標指標變數，把「它的位址」交給 As（所以是 &target）。
	var nfErr *NotFoundError
	if errors.As(err, &nfErr) {
		// 成功還原成 *NotFoundError，就能讀結構化欄位，而不是只拿到一坨字串。
		fmt.Printf("errors.As 挖到 NotFoundError → 資源=%s、ID=%s\n", nfErr.Resource, nfErr.ID)
	}

	// 對照：這條鏈跟付款無關，用 Is 比 ErrPaymentDeclined 會是 false。
	fmt.Println("這是付款被拒嗎？", errors.Is(err, ErrPaymentDeclined))

	// 一句話分工：sentinel + errors.Is 比「是不是某個固定錯誤」；
	// 自訂型別 + errors.As 拿「帶資料的錯誤細節」。
}
