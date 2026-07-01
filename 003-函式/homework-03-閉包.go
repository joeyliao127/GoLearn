//go:build ignore

// homework-03-閉包｜搭配 sample-03-閉包
// 目標（≤5 分鐘）：寫一個「工廠函式」makeDiscount，回傳一個會套用固定折扣的閉包。
// 電商情境：不同活動有不同折扣（VIP 打 8 折、清倉打 5 折…）。先做出一台「折扣機」，
// 之後把任何價格丟進去就自動打折。折扣用「折掉的百分比」表示，全程 int（別用 float）。
//
// TODO 1：宣告函式 makeDiscount(percent int) func(price int) int，
// 回傳一個閉包，它「捕捉」外層的 percent；閉包收 price、回傳折後價：
// 折後價 = price * (100 - percent) / 100（純整數運算）。
// TODO 2：在 main 用 vip := makeDiscount(20) 造一台「8 折機」，
// 再用 fmt.Println 印 vip(1000) 與 vip(500) 的結果。
//
// 完成後跑：go run 003-函式/homework-03-閉包.go
// 驗收：vip(1000)=800、vip(500)=400（8 折）；若改成 makeDiscount(50) 則 1000→500
package main

import "fmt"

// TODO 1：在這裡宣告 func makeDiscount(percent int) func(price int) int

func main() {
	// 在這裡開始寫 ↓（TODO 2：造一台 8 折機 vip，印 vip(1000) 與 vip(500)）
	fmt.Println("TODO: 還沒開始，把 TODO 1~2 完成後刪掉這行")
}
