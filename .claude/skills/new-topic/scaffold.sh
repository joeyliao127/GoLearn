#!/usr/bin/env bash
# scaffold.sh — 為 golearn 新主題建立「機械骨架」（由 new-topic skill 呼叫）
#
# 用法（在 repo 根目錄執行）：
#   bash .claude/skills/new-topic/scaffold.sh <NUM> <NAME> [子主題1 子主題2 ...]
# 例：
#   bash .claude/skills/new-topic/scaffold.sh 002 流程控制 if判斷 switch for迴圈 defer
#
# 會建立 NUM-NAME/，stamp 教學.html + 學習紀錄.html，並為每個子主題
# （依序給流水號 01,02,03…）stamp 一組 sample/homework 單檔（//go:build ignore）。
# 不傳子主題 → 只建資料夾與兩個 HTML（給 009/013 這種真 package 主題手動補 Go）。
#
# 只替換「機械佔位符」（NUM/NAME/FULLNAME/FILENAME/SAMPLENAME）；
# 教學內容（情境、章節、TODO、驗收）留給 Claude 依 課程規劃.md 填。
set -euo pipefail

SKILL_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TMPL="$SKILL_DIR/templates"
ROOT="$(pwd)"

NUM="${1:?需要主題編號，如 002}"
NAME="${2:?需要主題名，如 流程控制}"
shift 2 || true
FULLNAME="${NUM}-${NAME}"
DIR="$ROOT/$FULLNAME"

stamp() { # stamp <template> <target> [filename] [samplename]
  local tmpl="$1" target="$2" filename="${3:-}" samplename="${4:-}" c
  c="$(cat "$tmpl")"
  c="${c//\{\{NUM\}\}/$NUM}"
  c="${c//\{\{NAME\}\}/$NAME}"
  c="${c//\{\{FULLNAME\}\}/$FULLNAME}"
  c="${c//\{\{FILENAME\}\}/$filename}"
  c="${c//\{\{SAMPLENAME\}\}/$samplename}"
  printf '%s\n' "$c" > "$target"
}

if [ -e "$DIR" ]; then
  echo "⚠️  $DIR 已存在，未覆蓋。要重建請先刪掉或換編號。"; exit 1
fi

mkdir -p "$DIR"
stamp "$TMPL/教學.html.tmpl"     "$DIR/教學.html"
stamp "$TMPL/學習紀錄.html.tmpl" "$DIR/學習紀錄.html"
echo "✅ $FULLNAME/  （教學.html + 學習紀錄.html）"

n=0
for sub in "$@"; do
  n=$((n + 1)); NN="$(printf '%02d' "$n")"
  sfile="sample-${NN}-${sub}"
  hfile="homework-${NN}-${sub}"
  stamp "$TMPL/sample.go.tmpl"   "$DIR/${sfile}.go" "$sfile"
  stamp "$TMPL/homework.go.tmpl" "$DIR/${hfile}.go" "$hfile" "$sfile"
  echo "   + ${sfile}.go / ${hfile}.go"
done

echo ""
echo "骨架完成。下一步：依 課程規劃.md 與 001 範本，填入各檔教學內容，再 gofmt + go run 驗證。"
