//go:build ignore

// homework-03-sentinel與自訂｜搭配 sample-03-sentinel與自訂
// 目標（≤5 分鐘）：自己定義一個 sentinel error 並在查詢失敗時回傳它，呼叫方用 errors.Is 判斷。
// 電商情境（資料已給好）：用 SKU 查商品價格。查不到就回傳你定義的 sentinel。
//
// TODO 1：在套件層級（main 外面）宣告 sentinel：var ErrProductNotFound = errors.New("找不到商品")（慣例：名字 ErrXxx、用 errors.New）。
// TODO 2：在 priceOf 裡，若 map 查不到（用 v, ok := prices[sku] 的 ok 判斷），回傳 (0, ErrProductNotFound)；查得到則回傳 (價格, nil)。
// TODO 3：在 main 查 "NOPE"，用 errors.Is(err, ErrProductNotFound) 判斷；是就印「查無此商品」。
//
// 完成後跑：go run 008-錯誤處理/homework-03-sentinel與自訂.go
// 驗收：查 "NOPE" 時 errors.Is 為 true、印「查無此商品」；改查 "A100" 應印「價格 299 元」
package main

import (
	"errors"
	"fmt"
)

// TODO 1：在這裡宣告 var ErrProductNotFound = errors.New("找不到商品")

// 下面這行只是讓「還沒用到 errors 的半成品」能編譯；完成 TODO 後（errors 會被用到）就把這行刪掉。
var _ = errors.New

// priceOf：查得到回 (價格, nil)；查不到回 (0, ErrProductNotFound)。
func priceOf(sku string) (int, error) {
	prices := map[string]int{"A100": 299, "B200": 1290}

	// 下面這行 _ = ... 只是讓半成品能編譯；完成 TODO 2 用到 prices 後就把這行刪掉。
	_ = prices

	// TODO 2：用 v, ok := prices[sku] 判斷；ok 為 false 回 (0, ErrProductNotFound)，否則回 (v, nil)
	return 0, nil
}

func main() {
	price, err := priceOf("NOPE")

	// 下面兩行 _ = ... 只是讓半成品能編譯；完成 TODO 3 用到它們後就把這兩行刪掉。
	_ = price
	_ = err

	// 在這裡開始寫 ↓（TODO 3：errors.Is 判斷 ErrProductNotFound，或印出價格）
	fmt.Println("TODO: 還沒開始，把 TODO 1~3 完成後刪掉這行")
}
