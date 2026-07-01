// shipping_test.go｜對應「教學.html › 群組 A：1 testing 基礎、2 table-driven / t.Run」
// 跑法：go test ./013-testing/sample-01-basic （或加 -v 看每個子測試）
//
// 慣例三件事：
//  1. 檔名一定是 xxx_test.go（go test 只認這種檔名，正式編譯時會被忽略）。
//  2. 測試函式簽名固定 func TestXxx(t *testing.T)（Test 大寫開頭 + 收 *testing.T）。
//  3. 測試檔通常和被測程式「同 package」（這裡都是 package shipping），才測得到未匯出的東西。
package shipping

import "testing"

// ── 1. 最基本的測試：手動檢查一兩個值 ───────────────────────────
// 測試裡不用 assert 函式，就是「自己 if 比對，不符就 t.Errorf 回報」。
// 慣例訊息格式：Fee(輸入) = 實際, want 期望，之後看 log 就知道哪裡錯。
func TestFee(t *testing.T) {
	// 滿 1000 應免運。
	if got := Fee(1000); got != 0 {
		t.Errorf("Fee(1000) = %d, want %d", got, 0)
	}
	// 未滿 1000 收 60。
	if got := Fee(999); got != 60 {
		t.Errorf("Fee(999) = %d, want %d", got, 60)
	}
}

// ── 2. table-driven（表格驅動）：一個表 + 迴圈，一次涵蓋多案例 ──────
// Go 最常見的測試寫法：把「輸入/期望」列成一張表，再用迴圈跑。
// 每一列包一層 t.Run(子測試名, ...)，失敗時輸出會標出是哪個 case（如 Fee_Table/剛好滿千）。
func TestFee_Table(t *testing.T) {
	tests := []struct {
		name string // 子測試名稱，會顯示在輸出裡
		in   int    // 輸入：訂單金額
		want int    // 期望：運費
	}{
		{"剛好滿千免運", 1000, 0},
		{"超過門檻免運", 2500, 0},
		{"差一元未達門檻", 999, 60},
		{"金額為零收運費", 0, 60},
	}

	for _, tt := range tests {
		// t.Run 開一個子測試；即使某列失敗，其他列仍會繼續跑。
		t.Run(tt.name, func(t *testing.T) {
			if got := Fee(tt.in); got != tt.want {
				t.Errorf("Fee(%d) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}
