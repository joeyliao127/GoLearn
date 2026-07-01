//go:build ignore

// sample-02-命名慣例｜對應「教學.html › 4 節」
// 電商情境：用 Go 的命名慣例寫一段訂單邏輯。
// 跑法：go run 009-packages/sample-02-命名慣例.go
//
// 注意：這個子主題是單檔（跟命名有關、不牽涉跨包），所以用 `go run 檔.go`；
// 檔名可以是中文（單檔執行不經過 import 路徑）。
package main

import "fmt"

// 慣例：MixedCaps（駝峰），不用底線；常數也是 MixedCaps，不是 SCREAMING_CASE。
const maxItemsPerOrder = 100 // 小寫開頭 = 私有

// 縮寫整段大寫：ID / URL / HTTP（是 OrderID 不是 OrderId）。
type OrderID string

// 匯出、動詞開頭；小範圍參數用短名（unitPrice/qty 已夠清楚就不用更長）。
func CalcSubtotal(unitPrice, qty int) int {
	return unitPrice * qty
}

func main() {
	var id OrderID = "ORD-1001"
	fmt.Printf("訂單 %s 小計：%d 元（每單上限 %d 件）\n",
		id, CalcSubtotal(590, 3), maxItemsPerOrder)
}
