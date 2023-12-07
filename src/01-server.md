# 第一章&#x3000;簡單的 HTTP 伺服器

上網的時候，各位通常大概都發現過網站的網址最前面固定都會有一段 `http://` 或 `https://`。
我想，只有很少數的使用者思考過，這一段文字到底是什麼意思。

HTTP 是一種協議，也就是說，一種規定兩臺電腦之間的溝通方式的<a href="https://www.rfc-editor.org/rfc/rfc9110.html" target="_blank" rel="noopener noreferrer">文件</a>。
每一次想要透過 HTTP 協議得到一個資源，比如說當你打開此課程的時候，瀏覽器會在背後向某一臺伺服器「請求」每一個頁面。
對應每一個請求，伺服器將分別傳送一個「回應」。
提供資料的機器叫作伺服器（英：server），而請求資料的那一臺電腦則稱為客戶端（英：client）。

## 初始化專案

首先，我們來建立一個新專案。使用 `mkdir -p` 建立一個資料夾：

```shell
mkdir -p ~/projects/go-tutorial/01-server
```

`mkdir` 為 _make directory_ 的縮寫，而 `-p` 代表我們希望建立中間的資料夾。
然後進入稍早建立的資料夾並初始化新的專案 `go-tutorial/01-server`：

```shell
cd ~/projects/go-tutorial/01-server
go mod init go-tutorial/01-server
```

## 建立 Git 版本庫

開發任何軟體時，即使只是學習專案，使用一個版本控制系統是一個很好的慣例。
讓我們在學習專案的資料夾裡面建立一個 Git 版本庫：

```shell
cd ~/projects/go-tutorial
git init
```

Git 版本庫的變更單位稱作 commit，一個 commit 可以結合對於多個檔案的更動。
將 `~/projects/go-tutorial` 及所有子資料夾加到下一個 commit，借此告訴 Git，我們打算將所有更動加到下一個版本。
此動作的英文叫作 _stage for commit_，而對應的終端機指令為 `git add`。

此指令這後面不要忽略 `.`（一個半形句點），表示「當前資料夾」。加了當前資料夾也會一起加這個資料夾的所有子資料夾，包含稍早初始化的 `01-server`：

```shell
git add .
```

最後，儲存 commit：

```shell
git commit -m "Initial commit"
```

這是這個版本庫的第一個 commit，所以 _initial commit_ 就是「第一個版本」的意思。
每一個 commit 都必須包含一個 _commit message_，通常為這個版本80字符以內的英文描述。

若初次使用 Git，你的本地環境尚未設定 Git 信箱與姓名，這種情況下 Git 不會讓你完成 commit。請先設定 Git 的預設信箱與姓名：

```shell
git config --global user.name "Xiaoming Wang" # 填寫自己的資料
git config --global user.email "xiaoming.wang@example.com"
```

<!-- TODO: Get an example of an error message in a fresh environment -->

設定完本地 Git 版本庫，建議同步至 Github 當作備份與學習進度的紀錄。

## 第一個伺服器

在這個資料夾裡面，建立一個新檔案 `main.go`：

```go
{{#include ../code/01-server/main.go}}
```

### 程式碼分析

我們來分析這段程式碼。首先，這個檔案宣告 `package main`。在一個 Go 程式裏面，Go 都會先執行 `package main` 裡面的 `main` 函數：

```go
{{#include ../code/01-server/main.go:8:11}}
```

這一行將在螢幕上列出 `Listening on :3000...` 這段文字給人類看：

```go
{{#include ../code/01-server/main.go:9}}
```

至於為什麼要用 `log.Println` 而不是 `fmt.Println`，那就是因為 `log.Println` 除了文字，還會顯示當下的時間與日期，如下：

```plain
2023/12/06 00:54:30 Listening on :3000...
```

這一行反而比較複雜：

```go
{{#include ../code/01-server/main.go:10}}
```

`http.ListenAndServe` 將試圖在本地的 3000 號端口（英：port）啟動一個 HTTP 伺服器。`http.ListenAndServe` 的宣告如下：

```go
package http // import "net/http"

func ListenAndServe(addr string, handler Handler) error
```

這個函數的第一個參數為要希望監聽的端口，而第二個函數為一個 `http.Handler`。這個函數將返回一個錯誤。由於將這個函數的返回值傳達給 `log.Fatal`，如果監聽的動作發生問題，`log.Fatal` 就會印出錯誤並結束這段程式。
如果沒有出錯，那麼 `http.ListenAndServe` 將持續監聽我們所指定的端口並使用 `handler` 處理 HTTP 請求。

### `http.Handler` 的用法

至於 `http.ListenAndServe` 的第二個參數，`handler`，它的類型為 `http.Handler`。
`http.Handler` 是一個介面類型（英：interface），它的方法簽名要求一個方法：

```go
package http // import "net/http"

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

### `http.HandlerFunc` 的用法

`http` 這個 `package` 內另外提供了一種類型，名為 `http.HandlerFunc`：

```go
package http // import "net/http"

type HandlerFunc func(ResponseWriter, *Request)
    The HandlerFunc type is an adapter to allow the use of ordinary functions
    as HTTP handlers. If f is a function with the appropriate signature,
    HandlerFunc(f) is a Handler that calls f.

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
```

`http.HandlerFunc` 是最簡單的拿一般函數來處理 HTTP 請求的技巧。
只要我們宣告一個接受正確參數的函數：`func(http.ResponseWriter, *http.Request)`，就可以將它的類型轉換成 `http.HandlerFunc`：

```go
{{#include ../code/01-server/main.go:13:15}}
```

這個函數將 `<h1>Hello from HandleRequest!</h1>` 這段文字寫到一個 HTTP 請求的回應。然後，這樣定義的函數就可以當作 `http.HandlerFunc`：

```go
http.HandlerFunc(HandleRequest)
```

這樣的表達式就會符合 `http.Handler` 的要求，因此可以放在 `http.ListenAndServe` 第二個參數的位置：

```go
{{#include ../code/01-server/main.go:10}}
```

## 執行這段程式碼

現在，我們可以試著在專案的資料夾執行這段程式。如果輸入無誤，程式應該會編譯好，然後顯示 `Listening on :3000...`

```shell
$ go run .
2023/12/06 00:54:30 Listening on :3000...
```

打開瀏覽器，瀏覽至 `http://localhost:3000` 或 `http://127.0.0.1:3000` 應該就會看到以下畫面：

<figure class="bordered-figure">
<img src="/images/01/01.png" />
<caption>執行以上程式，瀏覽至<code>http://localhost:3000</code>的畫面。</caption>
</figure>

雖然我們沒有特別告知瀏覽器我們的內容將會返回 HTML，但由於我們所返回的內容**看起來**像 HTML，因此三大瀏覽器通通都將這段文字理解成 HTML。
