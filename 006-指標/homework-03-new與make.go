//go:build ignore

// homework-03-new與make｜搭配 sample-03-new與make
// 目標（≤5 分鐘）：用 make 建一個庫存表（map），寫入兩筆再讀出來。
// 電商情境：後台要初始化一份「商品名 → 庫存數」的表，先放兩個商品。
// 重點：map 一定要先 make（或用字面值）才能寫入，否則對 nil map 寫入會 panic。
//
// TODO 1：用 make 建立 stock，型別是 map[string]int（商品名 → 數量）
// TODO 2：寫入 stock["鍵盤"] = 10、stock["滑鼠"] = 25
// TODO 3：用 fmt.Printf 印出「鍵盤庫存：X，滑鼠庫存：Y」
//
// 完成後跑：go run 006-指標/homework-03-new與make.go
// 驗收：印出「鍵盤庫存：10，滑鼠庫存：25」
// （若把 make 那行拿掉、直接寫入 nil map，會 panic: assignment to entry in nil map）
package main

import "fmt"

func main() {
	// 在這裡開始寫 ↓
	// TODO 1：stock := make(map[string]int)
	// TODO 2：寫入兩筆庫存
	// TODO 3：印出兩筆庫存
	fmt.Println("TODO: 還沒開始，把 TODO 1~3 完成後刪掉這行")
}
