// Package signup｜練習搭配：這支程式「已經幫你寫好、可運作」，你的任務是幫它寫 mock 測試。
// 電商情境：使用者註冊時，把帳號存進某個 Repo（資料庫）。存去哪抽成介面，測試才能換成假的。
package signup

// Repo 是「存帳號」的抽象。任何有 Save(string) error 的型別都算實作它。
type Repo interface {
	Save(id string) error
}

// Register 是被測對象：把 id 交給注入的 Repo 存起來。
func Register(r Repo, id string) error {
	return r.Save(id)
}
