//go:build ignore

// sample-01-goroutine｜對應「教學.html › 1~3 節」
// 電商情境：一筆訂單成立後，要「同時」寄好幾封通知（Email、SMS、App 推播…）。
// 這些通知彼此獨立、又都在等網路 I/O，很適合各開一個 goroutine 並行處理，
// 主流程用 sync.WaitGroup 等「全部寄完」才收尾。
// 重點：① go 開一個 goroutine（幾乎零成本）；② WaitGroup 的 Add / Done / Wait；
// ③ goroutine 是非同步的，輸出順序不保證 —— 靠 Wait 收斂，別靠順序。
// 跑法：go run 010-併發基礎/sample-01-goroutine.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 要通知的管道（每個都當成一件耗時的 I/O 工作）。
	notifyChannels := []string{"Email", "SMS", "App 推播"}
	orderID := 1001

	// ── 1. goroutine：在函式呼叫前加 go，就把它丟到背景並行執行 ──────
	// 注意：如果 main 不等它們，main 自己跑完就結束了，背景還沒做完的
	// goroutine 會被一起收掉 —— 所以下面一定要用 WaitGroup 等。
	//
	// ── 2. sync.WaitGroup：等一組 goroutine 全部完成 ────────────────
	// 心法：開幾個就 Add 幾個（Add 要在 go 之前）；每個 goroutine 收尾 Done；主流程 Wait。
	var wg sync.WaitGroup

	for _, ch := range notifyChannels {
		wg.Add(1) // 每開一個 goroutine，計數 +1（務必在 go 之前 Add）

		// go func(...) 把這段丟到背景執行。
		// 慣例：把迴圈變數當「參數」傳進去（這裡的 ch），語意最清楚。
		// （Go 1.22 起迴圈變數每輪各自獨立，就算不傳也不會全部抓到最後一個；
		//   但顯式傳參數仍是最不易出錯、最好讀的寫法。）
		go func(channel string) {
			defer wg.Done() // ── 3. defer Done：中途 return 或 panic，計數一定會 -1 ──

			// 模擬寄送耗時（真實情況是打第三方 API）。
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("已寄出訂單 %d 的 %s 通知\n", orderID, channel)
		}(ch)
	}

	// Wait 會卡住，直到計數歸零（三個 goroutine 都 Done）才往下走。
	wg.Wait()

	// 這行保證在「全部通知都寄完」之後才印 —— 靠 Wait 收斂，而不是靠輸出順序。
	fmt.Printf("訂單 %d 的所有通知都寄送完成\n", orderID)
}
