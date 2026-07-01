// homework-02-mock｜練習：自己寫一個 mock 來測 Register（≤5 分鐘）
// 跑法：go test ./013-testing/homework-02-mock  （或加 -v）
//
// 目標：比照 sample-02-mock 的 mockNotifier，寫一個 fakeRepo 來測 Register。
// 現在是空測試（能通過但沒驗到東西），請照 TODO 補完。
package signup

import "testing"

func TestRegister(t *testing.T) {
	// TODO：
	// 1. 定義 type fakeRepo struct{ ... } 記錄「Save 被呼叫幾次、收到的 id」，
	//    並幫它寫 func (f *fakeRepo) Save(id string) error（實作 Repo 介面）。
	// 2. 建一個 &fakeRepo{}，呼叫 Register(fake, "user-1")，檢查回傳的 err 是否為 nil。
	// 3. 用 t.Errorf 斷言：Save 有被呼叫到、而且收到的 id == "user-1"。
	//
	// 提示：目前這是空測試，會直接 PASS；補完上面步驟後才真的有測到 Register。
}
