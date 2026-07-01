// sample-01-crosspkg（跨包）｜對應「教學.html › 1~2 節」
// 電商情境：main 跨 package 呼叫 pricing.Total 算含稅金額。
// 跑法：go run ./009-packages/sample-01-crosspkg
//
// 注意：這個主題是「真 package」示範，不是單檔——用 `go run ./資料夾`
// （帶點的路徑、對整個 package），不是 `go run 某個.go`。
// 也注意：資料夾名一定要用 ASCII（如 009-packages），因為它會變成 import
// 路徑的一部分，而 Go 的 import 路徑不接受中文（會 malformed import path）。
package main

import (
	"fmt"

	// import 路徑 = module 名(golearn) + 資料夾路徑（全 ASCII）。
	"golearn/009-packages/sample-01-crosspkg/pricing"
)

func main() {
	// 跨包呼叫用「套件名.函式名」。pricing 匯出的是大寫的 Total。
	fmt.Printf("含稅總額：%d 元\n", pricing.Total(590, 3))

	// 下面這行取消註解會編譯失敗：applyTax 小寫未匯出，跨包看不到。
	// fmt.Println(pricing.applyTax(100)) // ❌ cannot refer to unexported name
}
