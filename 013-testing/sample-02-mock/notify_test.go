// notify_test.go｜對應「教學.html › 群組 B：3 用介面注入 mock 來測試」
// 跑法：go test ./013-testing/sample-02-mock  （或加 -v）
//
// 重點：要測 Welcome，但不想真的寄信。做法＝寫一個「假的 Notifier」(mockNotifier)，
// 它實作 Send，但不真送、只把「被呼叫幾次 / 收到什麼訊息」記下來，最後用 t.Errorf 斷言。
// 這就是「用介面把依賴換成假實作來測」——mock 靠的還是「實作介面」。
package notify

import "testing"

// mockNotifier 是「為測試而生」的假 Notifier：不真送，只記錄互動。
// 它有 Send 方法 → 自動實作 Notifier，可以注入 Welcome。
type mockNotifier struct {
	calls   int    // Send 被呼叫幾次
	lastMsg string // 最後收到的訊息（驗證內容對不對）
}

func (m *mockNotifier) Send(msg string) error {
	m.calls++
	m.lastMsg = msg
	return nil // 假實作：永遠成功
}

func TestWelcome(t *testing.T) {
	mock := &mockNotifier{}

	// 注入 mock（而不是真的寄信器），呼叫被測函式。
	if err := Welcome(mock, "小明"); err != nil {
		// 這種「前置條件不成立就沒必要往下測」的情況，用 Fatalf 直接中止本測試。
		t.Fatalf("Welcome 回傳非預期錯誤：%v", err)
	}

	// 斷言 ①：Send 剛好被呼叫 1 次。
	if mock.calls != 1 {
		t.Errorf("Send 被呼叫 %d 次, want %d", mock.calls, 1)
	}
	// 斷言 ②：送出的訊息正確。
	if want := "歡迎 小明"; mock.lastMsg != want {
		t.Errorf("Send 收到訊息 = %q, want %q", mock.lastMsg, want)
	}
}
