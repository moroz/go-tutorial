# 如何使用本書

學習開發軟需要很多手動練習。
這一點對眾多曾接受填鴨式教育的華人讀者而言也許比較難理解。
開發軟體是一種實踐能力，沒有任何認證能夠證明你真的掌握了這門技術。
最後，面試第一份軟體開發工作的時候，最關鍵的是你曾經開發過的專案以及回答面試官的問題。

這些對於不想花錢的讀者而言也許是好消息：想要轉職不用花錢考證照，也沒有考試的時間壓力，唯獨需要投資你的時間與努力。
同時，如果你本來抱著考一個證照就可以轉職成功的心，這個課程也許不太適合你：我認為學習電腦一定要徹底了解電腦的操作方法及其背後的技術，而不只是背冷知識。

我曾經聽說一個台灣人做線上資訊科技課程 CS50 的時候根本沒有打開編輯器寫程式，反而只是看看影片，想要「儘快全部做完」。
CS50 是很好的課程，如果各位有空以及英文能力足夠，我強烈推薦認真做 CS50。
但 CS50 收穫最多的是它實踐的部分；光是第一個禮拜的功課就已經介紹了不少編程的基本觀念，而且還是用 C 語言。
如果認真折磨自己，用 C 語言寫出前幾個禮拜的作業，就會明白指標、函數、編譯、返回值等概念。
反而，如果只有走馬觀花地看一遍影片，恐怕什麼都學不到。

本書的目標是希望讓各位讀者初步了解網頁應用開發的幾項基本功：HTTP 協議、資料庫設計與 SQL 語言、HTML、CSS。

如果在文中出現等寬字體一段程式碼，請你打開你的編輯器，抄著範例打出這段程式碼並確認每一個小專案是否可以正確執行：

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

除非另有標示，可以假設程式碼都是 Go 語言。

如果文中出現 `$` 開頭的一段等寬文字，表示應該在終端機中執行的指令。隨之而來的文字為該指令的輸出：

```shell
$ go doc database/sql.DB.Query
package sql // import "database/sql"

func (db *DB) Query(query string, args ...any) (*Rows, error)
    Query executes a query that returns rows, typically a SELECT. The args are
    for any placeholder parameters in the query.

    Query uses context.Background internally; to specify the context, use
    QueryContext.
```

有時候需要執行多個指令，這時候我會省略開頭的 `$` 符號，這樣指令可以直接全部貼上：

```shell
mkdir -p ~/projects
cd ~/projects
```
