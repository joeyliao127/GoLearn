//go:build ignore

// homework-02-包裝與Is｜搭配 sample-02-包裝與Is
// 目標（≤5 分鐘）：把底層錯誤用 %w 包一層加上下文，再用 errors.Is 從包裝後的錯誤把它認出來。
// 電商情境（大半已給好）：套用折價券。低層 applyCoupon 發現券過期 → 回 sentinel ErrCouponExpired；
// 你要在 checkout 把它包上「結帳失敗」的脈絡，最外層再用 errors.Is 判斷是不是過期券。
//
// TODO 1：在 checkout 裡，如果 applyCoupon 回傳非 nil 的 err，用 fmt.Errorf("結帳失敗: %w", err) 包一層再回傳（注意動詞是 %w，不是 %v）。
// TODO 2：在 main 用 errors.Is(err, ErrCouponExpired) 判斷；是過期券就印「折價券已過期，請換一張」。
//
// 完成後跑：go run 008-錯誤處理/homework-02-包裝與Is.go
// 驗收：印出的完整訊息含「結帳失敗: 折價券已過期」，且 errors.Is 判斷為 true（印出換券提示）
package main

import (
	"errors"
	"fmt"
)

var ErrCouponExpired = errors.New("折價券已過期")

// applyCoupon：已寫好。code=="SUMMER" 視為過期券，回傳 sentinel。
func applyCoupon(code string) error {
	if code == "SUMMER" {
		return ErrCouponExpired
	}
	return nil
}

// checkout：呼叫 applyCoupon；出錯時「包一層」再回傳。
func checkout(couponCode string) error {
	err := applyCoupon(couponCode)
	// TODO 1：如果 err != nil，改成 return fmt.Errorf("結帳失敗: %w", err)
	return err
}

func main() {
	err := checkout("SUMMER")
	fmt.Println("完整錯誤：", err)

	// 下面這行 _ = ... 只是讓半成品能編譯；寫完 TODO 2 用到 err 後就把這行刪掉。
	_ = err

	// 在這裡開始寫 ↓（TODO 2：用 errors.Is(err, ErrCouponExpired) 判斷並印提示）
	fmt.Println("TODO: 還沒開始，把 TODO 1~2 完成後刪掉這行")
}
