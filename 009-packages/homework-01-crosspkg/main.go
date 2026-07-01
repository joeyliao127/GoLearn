// homework-01-crosspkg（跨包）｜搭配 sample-01-crosspkg
// 目標（≤5 分鐘）：完成 coupon 這個 package，讓 main 跨包算出折價後金額。
// 電商情境：套用折價券，折後金額不可低於 0。
//
// TODO：打開 coupon/coupon.go，把 Apply 實作出來（用它的未匯出 helper clampZero）。
//
// 跑法：go run ./009-packages/homework-01-crosspkg
// 驗收：Apply(1000, 300) = 700；Apply(1000, 1500) = 0（夾住、不變負數）
package main

import (
	"fmt"

	"golearn/009-packages/homework-01-crosspkg/coupon"
)

func main() {
	fmt.Println("折後金額（期望 700）：", coupon.Apply(1000, 300))
	fmt.Println("折後金額（期望 0）：", coupon.Apply(1000, 1500))
}
