//go:build ignore

// sample-03-tag｜對應「教學.html › 7~9 節」
// 電商情境：把 Product model 轉成 API 要回的 JSON、再從 JSON 讀回來。
// 這時就需要 struct tag——在欄位後面用反引號標 `json:"..."`，控制序列化出來的鍵名。
// 重點：① tag 語法（反引號字串，接在欄位後）；② 只有「大寫欄位」會被序列化；
// ③ 常用選項 `json:"名字,omitempty"`（零值時省略）與 `json:"-"`（完全不輸出）。
// （tag 只是「附註字串」，本身不做事——是 encoding/json 這種套件去讀它。深入在主題 012。）
// 跑法：go run 005-struct與方法/sample-03-tag.go
package main

import (
	"encoding/json"
	"fmt"
)

// ── 7. struct tag 語法：欄位後面接一段反引號字串 ──────────
// 格式慣例是 `key:"value"`；json 這個 key 就是給 encoding/json 看的。
// 沒寫 tag 的話，JSON 鍵名預設就用「欄位名原樣」（大寫開頭，像 "ID"）。
type Product struct {
	ID    string `json:"id"`    // 序列化成 "id"（而非預設的 "ID"）
	Name  string `json:"name"`  // 序列化成 "name"
	Price int    `json:"price"` // 序列化成 "price"

	// ── 8. 常用選項：omitempty 與 "-" ─────────────────────
	Discount int    `json:"discount,omitempty"` // 值為 0（零值）時，整個欄位省略不輸出
	Internal string `json:"-"`                  // "-" = 永遠不進出 JSON（例如內部備註）

	// ── 9. 陷阱：小寫欄位（未匯出）不會被序列化 ────────────
	// encoding/json 看不到 package 外的私有欄位，直接跳過——連 tag 都不用寫（寫了 go vet 還會警告）。
	costPrice int // 小寫 c → JSON 完全看不到它（就算硬加 json tag 也一樣被忽略）
}

func main() {
	// ── 序列化（Marshal）：struct → JSON ─────────────────────
	p := Product{
		ID:        "P001",
		Name:      "機械鍵盤",
		Price:     2490,
		Discount:  0,     // 零值 + omitempty → 不會出現在 JSON
		Internal:  "限量款", // tag 是 "-" → 不會出現在 JSON
		costPrice: 1200,  // 小寫欄位 → 不會出現在 JSON
	}

	// MarshalIndent 跟 Marshal 一樣，只是多排版（縮排）方便閱讀。
	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println("序列化失敗：", err)
		return
	}
	fmt.Println("轉成 JSON（注意 discount / Internal / costPrice 都不見了）：")
	fmt.Println(string(b))

	// ── 反序列化（Unmarshal）：JSON → struct ────────────────
	// 傳入的 JSON 有 discount=100，這次就會被讀進 Discount 欄位。
	// 注意要傳「指標」&p2，Unmarshal 才改得到它（呼應主題 006 的指標）。
	data := `{"id":"P002","name":"無線滑鼠","price":590,"discount":100}`
	var p2 Product
	if err := json.Unmarshal([]byte(data), &p2); err != nil {
		fmt.Println("反序列化失敗：", err)
		return
	}
	fmt.Printf("\n從 JSON 讀回來：%+v\n", p2)
	fmt.Printf("讀到的折扣 = %d 元\n", p2.Discount)
}
