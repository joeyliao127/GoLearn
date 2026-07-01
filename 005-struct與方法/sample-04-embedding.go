//go:build ignore

// sample-04-embedding｜對應「教學.html › 10~12 節」
// 電商情境：很多 model（Order、Product、Customer…）都需要「建立時間 / 建立者」這種
// 稽核欄位。與其每個 struct 都複製貼上一份，不如抽成一個 Audit，用「嵌入(embedding)」
// 塞進去共用。這是 Go 版的「不要重複」，用「組合」達成、而不是繼承。
// 重點：① 匿名欄位 = 嵌入；② 嵌入後欄位/方法會「晉升」，可直接 order.CreatedAt 存取；
// ③ Go 沒有繼承，這是「組合(composition)」——Order 有一個 Audit，不是 Order 是一種 Audit。
// 跑法：go run 005-struct與方法/sample-04-embedding.go
package main

import "fmt"

// ── 10. 先做一個可重用的小 struct，帶欄位也帶方法 ─────────
// Audit 蒐集「稽核資訊」，並提供一個 Summary() 方法。
type Audit struct {
	CreatedBy string
	CreatedAt string // 這裡先用字串代表時間；真正的 time.Time 留到主題 012
}

func (a Audit) Summary() string {
	return fmt.Sprintf("由 %s 於 %s 建立", a.CreatedBy, a.CreatedAt)
}

// ── 11. 嵌入：把型別名「直接當欄位」寫進去，不給欄位名 ────
// 這行沒有欄位名、只有型別 Audit，就是「嵌入(embedded field)」。
// 對比 sample-01 的 `Product Product`（具名欄位、has-a），這裡是把 Audit 的內容「攤平」進 Order。
type Order struct {
	Audit // ← 匿名（嵌入）：Audit 的欄位與方法會「晉升」到 Order

	ID    string
	Total int
}

func main() {
	// 建立時，嵌入的部分要用「型別名當 key」填：Audit: Audit{...}。
	order := Order{
		Audit: Audit{CreatedBy: "system", CreatedAt: "2026-07-01"},
		ID:    "O-20260701-003",
		Total: 4980,
	}

	// ── 12. 欄位/方法「晉升」：可直接透過外層存取，像自己的一樣 ──
	// 明明 CreatedAt 定義在 Audit 裡，卻能直接寫 order.CreatedAt——這就是「晉升(promotion)」。
	fmt.Printf("訂單 %s，金額 %d 元\n", order.ID, order.Total)
	fmt.Printf("建立時間（晉升欄位，直接 order.CreatedAt）：%s\n", order.CreatedAt)
	fmt.Printf("稽核摘要（晉升方法，直接 order.Summary()）：%s\n", order.Summary())

	// 需要時，仍可用「完整路徑」存取內層本體：order.Audit.CreatedBy。
	// 兩種寫法指到同一個東西。
	fmt.Printf("完整路徑存取：order.Audit.CreatedBy = %s（＝ order.CreatedBy = %s）\n",
		order.Audit.CreatedBy, order.CreatedBy)

	// 也能單獨把嵌入的那塊拿出來用（它就是個一般的 Audit 值）。
	fmt.Printf("把嵌入的 Audit 整個取出：%+v\n", order.Audit)
}
