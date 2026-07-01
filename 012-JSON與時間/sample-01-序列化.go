//go:build ignore

// sample-01-序列化｜對應「教學.html › 1~4 節」
// 電商情境：一筆訂單（Order）要送給前端／別的服務，得先「序列化」成 JSON 字串（Marshal）；
// 收到別人傳來的 JSON 也要「反序列化」還原成 struct（Unmarshal）。這是 API 每天在做的事。
// 重點：① 只有「大寫開頭（exported）」的欄位才會被序列化；② 用 struct tag 控制 JSON 的鍵名；
// ③ omitempty 讓「零值」欄位不要出現在輸出；④ Unmarshal 把 JSON 塞回 struct。
// 跑法：go run 012-JSON與時間/sample-01-序列化.go
package main

import (
	"encoding/json"
	"fmt"
)

// ── 1~3 節：用 struct tag 定義 JSON 的長相 ─────────────────────
// Order 的每個欄位後面用「反引號」寫 struct tag：`json:"鍵名,選項"`。
// - 欄位名必須大寫開頭（exported），小寫的 encoding/json 看不到、不會輸出（見 internal 欄位）。
// - tag 決定 JSON 裡的鍵名（沒寫 tag 就用欄位名本身，大小寫照舊）。
// - omitempty：這個欄位是「零值」（0、""、nil、false）時，就從輸出裡拿掉。
type Order struct {
	ID       string `json:"id"`               // 一定會輸出，鍵名為 "id"
	Customer string `json:"customer"`         // 客戶名
	Amount   int    `json:"amount"`           // 金額用 int 存（元），別用 float 避免精度問題
	Coupon   string `json:"coupon,omitempty"` // 有優惠券才輸出；空字串就整個消失
	Note     string `json:"-"`                // tag 寫 "-" ＝ 明確「永不序列化」
	internal string // 小寫開頭：unexported，json 完全看不到（連 "-" 都不用寫）
}

func main() {
	// ── 4 節：Marshal —— struct → JSON ───────────────────────────
	// 這筆有優惠券，Coupon 不是零值 → 會出現在 JSON 裡。
	o := Order{
		ID:       "A-1001",
		Customer: "王小明",
		Amount:   1290,
		Coupon:   "SUMMER10",
		Note:     "這行不會被序列化",
		internal: "小寫欄位也不會",
	}

	// json.Marshal 回傳 ([]byte, error)。慣例：一定要檢查 err（雖然這裡幾乎不會錯）。
	b, err := json.Marshal(o)
	if err != nil {
		fmt.Println("marshal 失敗：", err)
		return
	}
	// []byte 印出來要轉成 string 才是人看得懂的文字。
	fmt.Println("Marshal 緊湊版：", string(b))

	// MarshalIndent(值, 前綴, 縮排) → 有排版的 JSON，方便 log／debug 時閱讀。
	pretty, _ := json.MarshalIndent(o, "", "  ")
	fmt.Println("MarshalIndent 排版版：")
	fmt.Println(string(pretty))

	// ── omitempty 對照：這筆沒有優惠券（Coupon 是零值 ""）───────────
	// 結果 JSON 裡「不會」有 coupon 這個鍵——這就是 omitempty 的效果。
	noCoupon := Order{ID: "A-1002", Customer: "李小華", Amount: 500}
	b2, _ := json.Marshal(noCoupon)
	fmt.Println("沒券（省略 coupon）：", string(b2))

	// ── 4 節：Unmarshal —— JSON → struct ─────────────────────────
	// 收到一段 JSON（來自前端或別的服務），把它塞回 Order。
	// 注意：Unmarshal 第二個參數要傳「指標」&back，它才改得到 back 的內容。
	incoming := []byte(`{"id":"A-2001","customer":"陳大同","amount":888,"coupon":"VIP5"}`)
	var back Order
	if err := json.Unmarshal(incoming, &back); err != nil {
		fmt.Println("unmarshal 失敗：", err)
		return
	}
	fmt.Printf("Unmarshal 還原：%+v\n", back)
	fmt.Printf("取單一欄位：金額 = %d 元\n", back.Amount)

	// 小提醒：JSON 裡多出來、struct 沒有的鍵，Unmarshal 會「安靜忽略」，不會報錯。
	// （想「完整接住未知欄位」的作法，留到 sample-03 的 map[string]any。）
}
