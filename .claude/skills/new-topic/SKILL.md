---
name: new-topic
description: 依「課程規劃.md」的 spec 建立 golearn 學習系統的一個新主題（sample/homework/教學.html/學習紀錄.html），比照 001-變數與型別 的結構與品質。用於產出下一個主題如 002-流程控制，或未來新增主題。
---

# new-topic — 建立一個學習主題

把一個主題（如 `002-流程控制`）從 `課程規劃.md` 的 spec 產出成完整教材，結構與品質**比照 `001-變數與型別/`**。

## 輸入
- 主題編號與名稱（args，如 `002-流程控制`）。
- `課程規劃.md`（根目錄）對應主題條目 = 內容依據。
- `CLAUDE.md` = 慣例（命名、`//go:build ignore`、5 分鐘作業、電商主題、四種 callout…）。
- `001-變數與型別/` = 品質與風格的活範本，照它做。

## 步驟
1. **讀** `課程規劃.md` 該主題條目（子主題、每組 sample/homework 情境與驗收、idiom/陷阱、教學章節、銜接）與 `CLAUDE.md`。
2. **判斷結構**：一般主題走「單檔模式」（下方 3~6）；**009 套件與慣例 / 013 測試** 走「真 package 分支」（見文末）。
3. **跑 scaffold（單檔模式）**，在 repo 根目錄執行：
   ```
   bash .claude/skills/new-topic/scaffold.sh <NUM> <NAME> <子主題1> <子主題2> ...
   ```
   子主題**依序 = 流水號 01,02,03…**，用簡短 slug（例：`if判斷 switch for迴圈 defer`）。會建立 `NNN-NAME/` + 教學.html + 學習紀錄.html + 每組 sample/homework 單檔骨架。
4. **填 Go 檔**：把每個 `sample-NN-*.go`/`homework-NN-*.go` 骨架填成完整內容——
   - **sample**：完整可跑、註解**逐段對應教學.html 章節**、電商情境。
   - **homework**：**未解** TODO 骨架，含電商情境 + 明確 TODO + 驗收數字/輸出，難度 **≤5 分鐘**。若半成品有未用變數，用 `_ = x` 讓它能編譯。
5. **填 教學.html**：依骨架裡的 `<!-- CLAUDE 填 … -->` 標記填 nav、mindmap、各 section（**每個 sample 一個 `.topic` 群組**，maplink 標「對應檔案 + 跑法」）、動手練習表、慣例小抄。callout 用 key/conv/trap/analogy；程式碼 `< > &` 要 escape。
6. **驗證（務必）**：
   - `gofmt -w NNN-NAME/*.go` → `gofmt -l NNN-NAME/*.go`（要空）。
   - 每個 `sample-*.go` `go run` 有正確輸出；每個 `homework-*.go` 能編譯（印 TODO 行）。
   - 教學.html 的內部連結與 `../學習總覽.html`、`學習紀錄.html` 可通。
7. **回報**：列出建立的檔案、驗證結果、每個 homework 的驗收數字。
   > **若你是被主流程叫來的 subagent：到此為止、不要動 `學習總覽.html`**（由主流程統一更新，避免多個並行 subagent 撞同一檔）。
   > 互動式單獨執行時，才順手更新 `學習總覽.html` 該主題列（狀態 / 連結 / 最後更新）。

## 填檔常見坑（避開）
- **gofmt 會重排縮排註解**：多行 `//` 註解的續行**不要用前導空白**（gofmt 1.19+ 會當成 preformatted code-block，插空 `//` 行 + tab 縮排）。續行齊左、或折成單行。填完務必 `gofmt -w` 再 `gofmt -l`（要空且 idempotent）。
- **清掉骨架殘留**：填完 `教學.html` 後，`grep -n 'CLAUDE' 教學.html` 應為空——不要留 `<!-- CLAUDE 填 … -->` 標記或頂端 scaffolding 指示語。
- **半成品 homework 要能編譯**：未用的 import 是**硬編譯錯**（不同於未用區域變數），必要時用 `var _ = pkg.Fn`（或 `_ = x`）佔位並註明「完成後刪」；會 panic 的示範用開關（如 `ready := false`）gate 住，讓未解狀態能乾淨 `go run` 印 TODO。

## 真 package 分支（僅 009、013，不套 `//go:build ignore`）
這兩個主題要「用資料夾當 package」跑（`go run ./dir` / `go test ./dir`），有一條**硬規則**：

> ⚠️ **資料夾名（及放 package 的子資料夾）必須 ASCII**。資料夾會變成 import 路徑的一部分，而 Go 的 import 路徑**不收非 ASCII**（中文會 `malformed import path: invalid char`，連 `go build ./...` 都會壞）。做法：用 `scaffold.sh` 產中文 folder 後 `mv` 成 ASCII（如 `009-套件與慣例` → `009-packages`），**HTML 內的標題/標籤仍保留中文**。純單檔子主題（`go run 檔.go`）的檔名仍可中文。

scaffold 只用來產兩個 HTML（**不傳子主題**），Go 部分手工建：
- **009（已完成，當範本）**：資料夾 `009-packages`；跨包 demo `sample-01-crosspkg/`（`main.go` + `pricing/`，匯出 `Total`、未匯出 `applyTax`），跑 `go run ./009-packages/sample-01-crosspkg`；另含單檔 `sample-02-命名慣例.go`；目錄結構（cmd/internal/pkg）併入教學。
- **013**：資料夾如 `013-testing`；真 package + `xxx_test.go`（`TestXxx`、table-driven、用 007 的介面 mock），跑 `go test ./013-testing`。

## 品質檢查點（對照 001）
sample 註解對得到教學章節、情境具體｜homework 真的 ≤5 分鐘且有驗收｜四種 callout 有用｜有慣例小抄與 mindmap｜gofmt 乾淨、sample 輸出正確。
