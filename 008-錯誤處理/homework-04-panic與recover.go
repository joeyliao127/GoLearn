//go:build ignore

// homework-04-panic與recover｜搭配 sample-04-panic與recover
// 目標（≤5 分鐘）：用 defer + recover 攔住一個 panic，讓程式印出提示後「正常結束」而不是崩潰。
// 電商情境（已給好）：safePrintLabel 要印運送標籤，但拿到 nil 訂單時內部會 panic。
// 你要在函式最上面用 defer 攔住 panic，把它變成一行提示，讓 main 能繼續往下跑。
//
// TODO 1：在 safePrintLabel「一開始」加一個 defer 匿名函式，裡面用 r := recover()；若 r != nil 就 fmt.Println("已攔截 panic：", r)（這樣函式就不會把程式弄崩）。
// TODO 2：做好 TODO 1 後，把 main 裡的 ready 改成 true，讓 safePrintLabel(nil) 真正被呼叫。
//
// 完成後跑：go run 008-錯誤處理/homework-04-panic與recover.go
// 驗收：不會整個程式崩潰；依序印「已攔截 panic： ...(nil 指標)...」與「主程式繼續執行」
package main

import "fmt"

type Order struct {
	ID string
}

func safePrintLabel(o *Order) {
	// TODO 1：在這裡用 defer func(){ if r := recover(); r != nil { ... } }() 攔住 panic

	// 這行在 o 為 nil 時會 panic（存取 nil 指標的欄位）——就是要被你 recover 接住的異常。
	fmt.Println("列印運送標籤，訂單：", o.ID)
}

func main() {
	// ready 是「安全開關」：還沒加好 recover 前先設 false，避免半成品直接被 panic 弄崩。
	// TODO 2：做好 safePrintLabel 的 recover 後，把這行改成 true。
	ready := false

	if !ready {
		fmt.Println("TODO: 還沒開始，先完成 safePrintLabel 的 recover，再把 ready 改成 true")
		return
	}

	safePrintLabel(nil) // 故意傳 nil 觸發 panic（recover 好了才輪到它跑）
	fmt.Println("主程式繼續執行")
}
