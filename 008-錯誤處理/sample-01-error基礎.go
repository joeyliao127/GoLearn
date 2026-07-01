//go:build ignore

// sample-01-error基礎｜對應「教學.html › 1~3 節」
// 電商情境：從錢包扣款（withdraw）。餘額不夠時，Go 不會 throw exception，
// 而是「把錯誤當成一個回傳值交回來」，呼叫方自己判斷。
// 重點：① Go 沒有 try/catch，error 就是普通的值（型別是內建的 error 介面）；
// ② 慣例是最後一個回傳值放 error：func f() (結果, error)；
// ③ 呼叫方用 if err != nil「早 return」，正常路徑往下走、不縮排。
// 跑法：go run 008-錯誤處理/sample-01-error基礎.go
package main

import (
	"errors"
	"fmt"
)

// ── 1. error 是「值」＋ (結果, error) 慣例 ──────────────────
// withdraw 從 balance 扣掉 amount，回傳 (剩餘餘額, error)。
// error 是內建介面，nil 代表「沒出事」；非 nil 代表「出事了，看這個訊息」。
func withdraw(balance, amount int) (int, error) {
	// 用 errors.New 做「固定字串」的錯誤；訊息小寫開頭、結尾不加標點（Go 慣例）。
	if amount <= 0 {
		return balance, errors.New("扣款金額必須大於 0")
	}
	// ── 2. fmt.Errorf：需要把「數字/變數」帶進訊息時用它 ──────
	// errors.New 只能吃固定字串；要格式化就用 fmt.Errorf（像 Printf 但回傳 error）。
	if amount > balance {
		return balance, fmt.Errorf("餘額不足：欲扣 %d 元，但只有 %d 元", amount, balance)
	}
	// 沒問題：回傳新餘額，error 位置給 nil。
	return balance - amount, nil
}

func main() {
	balance := 1000

	// ── 3. 呼叫方慣例：if err != nil 早 return，錯誤先處理掉 ───
	// Go 的招牌長相：拿到 (值, err)，先檢查 err，有問題就立刻收尾（這裡是印出並結束）。
	// 這叫「early return」——把錯誤路徑先擋掉，正常邏輯就能平鋪直敘、不用一直往右縮排。
	newBalance, err := withdraw(balance, 300)
	if err != nil {
		fmt.Println("扣款失敗：", err)
		return
	}
	fmt.Printf("扣款成功，餘額 %d 元\n", newBalance)

	// 這次故意扣超過餘額，讓 error 走非 nil 分支。
	// 注意：newBalance 已宣告過，所以這裡用 = 重新賦值，不能再用 :=。
	newBalance, err = withdraw(newBalance, 5000)
	if err != nil {
		// 錯誤是「值」，可以直接印、可以往上回傳、可以留著等等比對——它就是一般資料。
		fmt.Println("扣款失敗：", err)
		return
	}
	fmt.Printf("扣款成功，餘額 %d 元\n", newBalance)

	// 這行不會被執行到，因為上面 err != nil 已經 return 了。
	fmt.Println("（這行印不出來：上面已因餘額不足而 return）")
}
