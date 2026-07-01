//go:build ignore

// homework-01-if判斷｜搭配 sample-01-if判斷
// 目標（≤5 分鐘）：用 if / else if / else 依「會員等級」給折扣，算出應付金額。
// 電商情境（輸入已給好，直接判斷）：一筆原價 1000 元的訂單，會員等級 level。
// 等級規則：level>=3（鑽石）打 8 折；>=2（金卡）打 9 折；>=1（銀卡）打 95 折；其餘（一般）不打折。
//
// TODO 1：用 if / else if / else 判斷 level，算出折扣率 rate（float64）；8 折＝0.8、95 折＝0.95、不打折＝1.0
// TODO 2：算出應付金額 payable（int）＝ price × rate；price 是 int、rate 是 float64，先 float64(price)*rate 再 int(...) 轉回整數
// TODO 3：用 fmt.Printf 印出「等級 X → 應付 Y 元」
//
// 完成後跑：go run 002-流程控制/homework-01-if判斷.go
// 驗收：level=2（金卡）時應付 900 元；把 level 改成 3 應付 800、改成 0 應付 1000
package main

import "fmt"

func main() {
	price := 1000
	level := 2 // 想驗證別的等級，改這個數字（0~3）

	// 下面兩行 _ = ... 只是讓半成品能編譯；用到某個變數後就把對應那行刪掉。
	_ = price
	_ = level

	// 在這裡開始寫 ↓
	fmt.Println("TODO: 還沒開始，把 TODO 1~3 完成後刪掉這行")
}
