// homework-01-basic｜練習：幫 Apply 寫 table-driven 測試（≤5 分鐘）
// 跑法：go test ./013-testing/homework-01-basic  （或加 -v）
//
// 目標：把下面的 TestApply 改成 table-driven（比照 sample-01-basic 的 TestFee_Table）。
// 現在只放了一個起始 case，先讓 go test 能通過；請照 TODO 補齊其他案例。
package discount

import "testing"

func TestApply(t *testing.T) {
	// 起始案例：打 9 折。這個先讓測試能跑、能過。
	if got := Apply(1000, 10); got != 900 {
		t.Errorf("Apply(1000, 10) = %d, want %d", got, 900)
	}

	// TODO：用 table-driven 補更多案例（例如 Apply(1000,10)==900、Apply(1000,100)==0、
	// Apply(1000,0)==1000）。做法：宣告 []struct{name string; price, percent, want int}，
	// 用 for range + t.Run(tt.name, ...) 逐列檢查。可把上面那個起始 case 也搬進表格。
}
