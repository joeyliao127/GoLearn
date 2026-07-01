//go:build ignore

// homework-02-命名慣例｜搭配 sample-02-命名慣例
// 目標（≤5 分鐘）：把下面「不符 Go 慣例」的命名改成慣例寫法。
// 跑法：go run 009-packages/homework-02-命名慣例.go
//
// TODO：修正這 4 個命名（改完程式仍要能跑）：
//  1. const MAX_RETRY（SCREAMING_CASE）→ MixedCaps：maxRetry
//  2. type UserId → UserID（縮寫 ID 全大寫）
//  3. func get_user_name（蛇底線）→ MixedCaps：getUserName
//  4. 變數 http_url → httpURL（縮寫全大寫）
//
// 提示：Go 慣例＝MixedCaps、不用底線、縮寫整段大寫；大寫=匯出、小寫=私有。
package main

import "fmt"

const MAX_RETRY = 3

type UserId int

func get_user_name(id UserId) string { return fmt.Sprintf("會員-%d", id) }

func main() {
	http_url := "https://shop.example.com"
	// 這幾行 _ = ... 讓半成品能編譯；改好命名、真的用到它們後就刪掉。
	_ = MAX_RETRY
	_ = http_url
	fmt.Println(get_user_name(1))
	fmt.Println("TODO: 依上面 4 點改成 Go 命名慣例，改完刪掉這行")
}
