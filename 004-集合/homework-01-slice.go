//go:build ignore

// homework-01-slice｜搭配 sample-01-slice
// 目標（≤5 分鐘）：用 slice 建購物車、append 加購新商品、算出「總件數」與「總金額」。
// 電商情境：使用者又加了一件商品到購物車，要重算結帳頁的件數與小計。
//
// TODO 1：用 append 把 newItem 加進 cart（記得寫成 cart = append(cart, newItem)，回傳值要接回去）
// TODO 2：用 for range 走訪 cart，累加出 count（int，總件數）與 total（int，總金額＝各 Price 相加）
// TODO 3：用 fmt.Printf 印出「共 X 件，小計 Y 元」
//
// 完成後跑：go run 004-集合/homework-01-slice.go
// 驗收：加購後共 4 件、小計 4120 元（2490+890+390+350）
package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	cart := []Product{
		{Name: "機械鍵盤", Price: 2490},
		{Name: "無線滑鼠", Price: 890},
		{Name: "滑鼠墊", Price: 390},
	}
	newItem := Product{Name: "手機支架", Price: 350} // 要加購的商品

	// 下面兩行 _ = ... 只是讓半成品能編譯；開始用到某個變數後就把對應那行刪掉。
	_ = cart
	_ = newItem

	// 在這裡開始寫 ↓
	fmt.Println("TODO: 還沒開始，把 TODO 1~3 完成後刪掉這行")
}
