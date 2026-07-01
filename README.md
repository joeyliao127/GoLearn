# 🐹 golearn — 電商導向的 Go 實戰自學環境

> 一套為「**快速接手一個電商 Go 專案**」而設計的、系統化、動手做、有反饋迴路的 Go 學習環境。
> 每個主題都有 **可跑的範例 + 對應教學 + 5 分鐘小作業**，全部用電商情境（訂單／商品／購物車／庫存／金流／會員）。

適合：**想快速具備「讀懂並小幅貢獻電商 Go 專案」基礎能力的人**，尤其是已經接手（或即將接手）Go 後端、時間有限、需要用社群慣用方式上手的工程師。

---

## 為什麼這樣設計（設計前提）

這個 repo **不是完整的 Go 教程**，而是刻意收斂的「上手包」。設計上的取捨：

- **目標是「夠快速接手專案」的基礎**，不是學完整個 Go。用最短路徑覆蓋讀電商 repo 幾乎一定會撞到的東西。
- **全電商主題**：學到的每個觀念都能直接對應到你手上的專案。
- **每個作業 ≤ 5 分鐘**：降低啟動阻力、維持節奏；一個子主題若不夠驗收，就拆成多個小作業，而不是把單題變難。
- **有反饋迴路**：光「看」留不住，靠「輸出 → 被糾錯 → 間隔複習」才扎實。所以每個主題都配了作業、批改紀錄、弱項追蹤。
- **教材能獨立自學**：`sample` 的註解逐段對應 `教學.html` 的章節，讀範例不必兩邊切換。

---

## 環境需求

- **Go 1.26+**（本 module 使用 `go 1.26.4`；多數範例在較舊版本也能跑，但個別主題會用到新版行為，例如 Go 1.22 的迴圈變數語意）。
- 一個能開 HTML 的**瀏覽器**看教學文件（推薦 VSCode 內建 **Simple Browser**；教學是純前端、無建置、單一自包含 HTML，離線也讀得了，只有心智圖需要連網用 CDN 渲染）。
- **（選配）[Claude Code](https://claude.com/claude-code)** —— 用來自動批改作業、維護「學習紀錄」。**沒有它也能完整自學**（見下方）。

## 快速開始

```bash
git clone <this-repo> && cd golearn
go version                                    # 確認 Go 裝好

# 跑第一個範例（注意：單檔執行，見下方「怎麼跑範例」）
go run 001-變數與型別/sample-01-變數宣告.go

# 用瀏覽器開這個當儀表板入口：
#   學習總覽.html
```

---

## 怎麼用這個 repo 學（學習方法）

**以 repo 教材為主、依編號 `001 → 013` 順序學。** 每個主題跑一輪：

1. **讀** 該主題的 `教學.html`，同時對照 `sample-*.go`（主要教材，邊讀邊 `go run` 看輸出）。
2. **（選）** 自己找一支對應該主題的 YouTube 影片補強——影片是輔助，不是主軸。
3. **做** `homework-*.go`：照檔案裡的 `TODO` 完成（≤5 分鐘），`go run` 跑跑看，對照註解裡的**「驗收」數字**自我檢查。
4. **（選，需 AI）** 把作業交給 Claude 批改 → 它會把評分／可優化點／弱項寫進該主題的 `學習紀錄.html`。
5. **回** `學習總覽.html` 看進度與弱項 → 進入下一個主題。

### 沒有 AI 助手也能學

核心（`sample` + `教學.html` + `homework` + 每題的**驗收數字**）是**完全可獨立自學**的：讀教材 → 寫作業 → `go run` 跟驗收數字對照，就能自我檢查對錯。`學習紀錄.html` 這種「AI 自動批改與弱項追蹤」是加分項，不是必需。

### 怎麼跑範例（重要）

每個 `sample` / `homework` 都是**自成一格的獨立單檔**（檔頭有 `//go:build ignore`），用 `go run` **單檔執行**：

```bash
go run 002-流程控制/sample-04-defer.go
go run 008-錯誤處理/sample-02-包裝與Is.go
```

> ⚠️ 因為整個 repo 的檔案都是獨立腳本，`go build ./...` / `go vet ./...` 會顯示 **`matched no packages`**——這是**預期行為**，不是壞掉。要靜態檢查請用單檔：`go vet <檔>`；要檢查格式：`gofmt -l 資料夾/*.go`。
>
> （這也是刻意的教學設計：同一資料夾放多個 `main()`，靠 `//go:build ignore` 讓它們互不干擾、半成品作業也不會弄壞整包 build。）

---

## 反饋迴路（這個 repo 的核心機制）

```
sample（教）→ 教學.html（讀）→ homework（練，≤5 分鐘）
   → AI 批改（回饋）→ 學習紀錄.html（存 QA／錯誤／評分／弱項）
   → 學習總覽.html（追進度 + 弱項總表）
```

每個主題資料夾都有一份 `學習紀錄.html`（六區塊：本主題摘要、Homework 紀錄、QA 問答、踩坑與訂正、弱項雷達、待複習），根目錄的 `學習總覽.html` 則把各主題的進度與弱項彙整成一張儀表板。

---

## 課程地圖（13 主題）

編號即學習順序。目前狀態（教材是否就緒）：

| 階段 | 編號 | 主題 | 內容 | 狀態 |
|---|---|---|---|---|
| 基礎語法 | 001 | 變數與型別 | var/const/iota、基本型別+零值、型別轉換、運算符、scope | ✅ |
| 基礎語法 | 002 | 流程控制 | if/switch、for（Go 沒有 while）、defer | ✅ |
| 基礎語法 | 003 | 函式 | 多回傳值 `(值, error)`、variadic、閉包 | ✅ |
| 資料結構 | 004 | 集合 | slice/map/array、nil 行為、aliasing 陷阱 | ✅ |
| 資料結構 | 005 | struct 與方法 | struct、method/receiver、struct tag、組合 embedding | ✅ |
| 資料結構 | 006 | 指標 | 傳值心法、`*T`/`&x`/`*p`、pointer receiver、new/make | ✅ |
| 資料結構 | 007 | 介面 | interface、duck typing、用介面做 DI 與 mock、nil interface 陷阱 | ✅ |
| 工程實務 | 008 | 錯誤處理 | error 是值、`%w` 包裝、`errors.Is/As`、sentinel、panic/recover | ✅ |
| 工程實務 | 009 | 套件與慣例 | package/匯出、cmd/internal/pkg 目錄結構、命名 | 🚧 待建 |
| 工程實務 | 010 | 併發基礎 | goroutine、channel、WaitGroup、Mutex、select | ✅ |
| 工程實務 | 011 | context | 傳遞、逾時/取消、request-scoped value | ✅ |
| 工程實務 | 012 | JSON 與時間 | encoding/json、struct tag、time.Time | ✅ |
| 工程實務 | 013 | 測試 | testing、table-driven、用介面 mock | 🚧 待建 |

> **009／013** 需要「真正的 package／`*_test.go`」結構（跟其他主題的單檔模式不同），規劃中。
> **基礎之後的路線**（尚未排入）：HTTP 框架（gin/echo/chi）、資料庫（database/sql、sqlx、gorm）、generics。

---

## repo 結構

```
golearn/
├── 學習總覽.html            # 入口儀表板：課程路線圖 + 各主題進度 + 弱項總表
├── 課程規劃.md              # 每個主題的設計 spec（給擴充用）
├── CLAUDE.md               # 設計理念、慣例、給 AI 助手的教學指引
├── NNN-主題/                # 每個資料夾 = 一個主題（編號即順序）
│   ├── 教學.html            #   主要教材（章節對應 sample 註解）
│   ├── 學習紀錄.html         #   反饋紀錄（AI 維護）
│   ├── sample-NN-<子主題>.go #   完整可跑的教學範例
│   └── homework-NN-<子主題>.go #  未解的 5 分鐘小作業
├── .claude/skills/new-topic/ # 「新增主題」的可重用 skill（見下）
└── _封存/                    # 舊版學習內容（僅供參考，不進學習動線）
```

**檔名慣例**：`sample-NN-<子主題>.go` / `homework-NN-<子主題>.go`；`NN` 是資料夾內從 01 起的流水號（號碼在前方便排序），sample 與 homework 同號成對。

---

## 📌 資料來源與正確性

教材涉及的 **Go 語法與標準函式庫行為**，在製作時採「**雙保險**」查證，而非憑印象：

1. **[Context7](https://context7.com)** 取得的 **Go 官方文件**（`go.dev` 說明文件 + `golang/go` 原始碼）對照查證；
2. 每個主題的範例都實際 **`go run` 執行**、以 **`go vet`** 檢查、併發主題另跑 **`go run -race`** 確認無資料競爭。

- **參考時間：2026-07**（對應 **Go 1.26.4**）。
- 有特別查證過的點包括：`//go:build ignore` 檔用 `go run <檔>` 的執行行為、slice/map 的 nil 與 aliasing 語意、`nil` interface 陷阱（interface = (型別, 值)）、`errors.Is`/`errors.As`/`%w`、`context` 的 `WithTimeout`/`WithValue` 慣例、`time.Time` 的 RFC3339 序列化、Go 1.22 起的迴圈變數語意等。

> ⚠️ Go 會持續演進。若你的 Go 版本與上述不同，請以 `go doc`、[go.dev](https://go.dev) 官方文件與**實際執行結果**為準。教材以「當前社群通用寫法」為主（例如錯誤處理教 `errors.Is/As`，而非過新、尚未普及的變體）。

---

## 如何擴充（新增主題）

repo 內建一個可重用的 **`new-topic` skill**（`.claude/skills/new-topic/`），把「建一個主題」的流程、樣板與骨架腳本都封裝好：

- **用 [Claude Code](https://claude.com/claude-code)**：執行 `/new-topic 014-你的主題`，它會依 `課程規劃.md` 的 spec + 比照既有主題產出完整教材。
- **手動**：照 `SKILL.md` 的流程，先跑 `scaffold.sh` 產骨架，再依 `課程規劃.md` 填入內容。

```bash
# 產出某主題的機械骨架（資料夾 + 兩個 HTML + 每組 sample/homework 檔頭）
bash .claude/skills/new-topic/scaffold.sh 014 你的主題 子主題A 子主題B
```

---

## 給不同的人

- **想學 Go 基礎、快速上手電商專案** → 開 `學習總覽.html`，照 `001 → 013` 一路做作業。
- **想理解／接手這個學習專案本身** → 讀 `CLAUDE.md`（設計理念與慣例）+ `課程規劃.md`（課程 spec）。
- **想擴充主題或改造成別的領域** → 用 `new-topic` skill；把 `課程規劃.md` 與各主題的電商情境換成你的領域即可。

---

## 定位與免責

- 這是一個**學習練習場**，**不是 production code**：範例刻意簡化、聚焦單一觀念，為了教學清楚而非工程完備（例如金額一律用 `int` 存最小單位、示範用的資源清理是模擬的）。
- `_封存/` 是重建前的舊學習內容，保留供參考，不在學習動線內。
- 繁體中文為主。自由 fork／取用／改造。

---

<sub>🐹 用 Go 官方色 `#00ADD8` 打造 · 教材以 context7 對照官方文件（2026-07）並實跑驗證 · 動手優先、慢慢累積</sub>
