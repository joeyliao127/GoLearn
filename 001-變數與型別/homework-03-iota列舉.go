//go:build ignore

// homework-03-iota列舉｜搭配 sample-03-iota列舉
// 目標（≤5 分鐘）：用 iota 做一組「會員等級」列舉，並讓它印出中文。
// 電商情境：會員分四級，等級越高折扣越好。
//
// TODO 1：宣告一個新型別 type MemberLevel int
// TODO 2：用 const + iota 定義四級（0 起自動編號）：Normal 一般 = iota、Silver 銀卡、Gold 金卡、Diamond 鑽石卡
// TODO 3：幫 MemberLevel 加 String() 方法，回傳中文（一般 / 銀卡 / 金卡 / 鑽石卡）
// TODO 4：在 main 裡 level := Gold，用 %v 印「你的等級：金卡」，再用 %d 印底層數字（應為 2）
//
// 提示：整組寫法可完全對照 sample-03-iota列舉 的 OrderStatus。
// 完成後跑：go run 001-變數與型別/homework-03-iota列舉.go
// 驗收：印出「你的等級：金卡」與底層數字 2
package main

import "fmt"

// TODO 1 & 2：在這裡宣告 MemberLevel 型別與四個等級常數

// TODO 3：在這裡幫 MemberLevel 加 String() 方法

func main() {
	// TODO 4：宣告 level := Gold，用 %v 和 %d 印出來
	fmt.Println("TODO: 還沒開始，把 TODO 1~4 完成後刪掉這行")
}
