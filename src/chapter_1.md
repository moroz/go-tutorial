# 第一章 簡單的 HTTP 伺服器

上網的時候，各位通常大概都發現過網站的網址最前面固定都會有一段 `http://` 或 `https://`。
我想，只有很少數的使用者思考過，這一段文字到底是什麼意思。

HTTP 是一種協議，也就是說，一種規定兩臺電腦之間的溝通方式的<a href="https://www.rfc-editor.org/rfc/rfc9110.html" target="_blank" rel="noopener noreferrer">文件</a>。
每一次想要透過 HTTP 協議得到一個資源，比如說當你打開此課程的時候，瀏覽器會在背後向某一臺伺服器「請求」每一個頁面。
對應每一個請求，伺服器將分別傳送一個「回應」。
提供資料的機器叫作伺服器（英：server），而請求資料的那一臺電腦則稱為客戶端（英：client）。

尚未接觸 HTTP 協議的同學可能會對

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

在這個資料夾裡面，建立一個新檔案 `main.go`：

```go
{{#include ../code/01-server/main.go}}
```

我們來分析這段程式碼。首先，這個檔案宣告 `package main`。在一個 Go 程式裏面，Go 都會先執行 `package main` 裡面的 `main` 函數：

```go
{{#include ../code/01-server/main.go:8:11}}
```

這一行將在螢幕上列出 `Listening on :3000...` 這段文字給人類看：

```go
{{#include ../code/01-server/main.go:9}}
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

至於 `http.ListenAndServe` 的第二個參數，`handler`，它的類型為 `http.Handler`：

```go
package http // import "net/http"

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

`http.Handler` 為一個介面（英：interface），它要求

```go
{{#include ../code/01-server/main.go:13:15}}
```

<figure class="bordered-figure">
<img src="/images/01/01.png" />
<caption>執行以上程式，瀏覽至<code>http://localhost:3000</code>的畫面。</caption>
</figure>
