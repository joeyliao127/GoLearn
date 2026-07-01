//go:build ignore

// sample-01-struct｜對應「教學.html › 1~3 節」
// 電商情境：定義電商最核心的兩個 model——Product（商品）與 Order（訂單），
// 並示範幾種「建立 struct 值」的寫法。這就是你在電商 repo 裡最常看到的資料結構。
// 重點：① struct 是「把一組欄位綁在一起」的自訂型別；② 建立時用「欄位名: 值」的
// 具名寫法最清楚；③ &T{} 直接拿到指標、new(T) 給零值指標——兩者差異這裡先點到。
// 跑法：go run 005-struct與方法/sample-01-struct.go
package main

import "fmt"

// ── 1. 定義 struct：把相關欄位綁成一個型別 ────────────────
// 欄位大寫開頭 = 匯出（別的 package 可存取），小寫 = 私有。電商 model 的欄位
// 幾乎都大寫，因為之後常要跨 package 用、或被 JSON 序列化（見 sample-03）。
// 金額一律用 int 存「元」，別用 float64——浮點數算錢會有精度誤差。
type Product struct {
	ID    string
	Name  string
	Price int // 單價（元）
	Stock int // 庫存數量
}

// Order 內含另一個 struct（Product）當欄位，這叫「組合」——電商 model 常這樣層層包起來。
type Order struct {
	ID       string
	Customer string
	Product  Product // 一個 Order「有一個」Product（has-a 關係）
	Qty      int
}

func main() {
	// ── 2. 建立 struct 值：三種寫法 ──────────────────────────
	// (a) 具名欄位（推薦）：寫出欄位名，順序可隨意、沒填的欄位自動補「零值」。
	//     這是電商 code 的主流寫法，加/改欄位也不會壞、可讀性最好。
	keyboard := Product{
		ID:    "P001",
		Name:  "機械鍵盤",
		Price: 2490,
		Stock: 50,
	}

	// (b) 位置寫法（不推薦）：不寫欄位名、純靠順序填。欄位一多就難讀、加欄位就爆。
	//     這裡示範一下長相就好，實務上請用具名。
	mouse := Product{"P002", "無線滑鼠", 590, 120}

	// (c) 只填部分欄位：沒填的自動是零值（string="", int=0, bool=false）。
	//     這裡故意不填 Stock，看它自動變 0。
	cable := Product{ID: "P003", Name: "USB-C 線", Price: 190}

	fmt.Printf("鍵盤：%+v\n", keyboard) // %+v 會連「欄位名」一起印，debug 神器
	fmt.Printf("滑鼠：%+v\n", mouse)
	fmt.Printf("線材（Stock 未填，自動為 0）：%+v\n", cable)

	// ── 3. 存取 / 修改欄位，以及 &T{} vs new(T) ──────────────
	// 用「點」存取欄位；struct 是可變的，直接改欄位即可。
	keyboard.Stock -= 1 // 賣掉一個，庫存 -1
	fmt.Printf("賣出一個鍵盤後庫存：%d\n", keyboard.Stock)

	// 建一張訂單，把 keyboard 塞進去（注意：這裡是「複製一份」進 Order，
	// 之後改 order.Product 不會動到外面的 keyboard——值型別的複製語意，主題 006 會講透）。
	order := Order{
		ID:       "O-20260701-001",
		Customer: "王小明",
		Product:  keyboard,
		Qty:      2,
	}
	fmt.Printf("訂單：%s 買了 %d 個 %s\n", order.Customer, order.Qty, order.Product.Name)

	// &T{}：建立 struct 並「直接拿到它的指標」（*Product）。電商 code 建 model
	// 常回傳指標（省一次複製、也方便讓方法改到本體，見 sample-02 與主題 006）。
	pKeyboard := &Product{ID: "P001", Name: "機械鍵盤", Price: 2490, Stock: 50}
	// 有了指標，Go 貼心地讓你「照樣用點存取」，不用先寫 (*pKeyboard).Name。
	fmt.Printf("用 &T{} 拿到指標：%T，商品名 = %s\n", pKeyboard, pKeyboard.Name)

	// new(T)：配置一個「全零值」的 T 並回傳指標，等同 &T{} 但不能順便填欄位。
	// 想建立時就帶初始值 → 用 &T{}；只想要個零值空殼再慢慢填 → new(T) 也行。
	blank := new(Product) // 等同 &Product{}
	blank.Name = "待補商品"
	fmt.Printf("用 new(T) 拿到零值指標再填：%+v\n", *blank)
}
