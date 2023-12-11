# 第一章&#x3000;簡單的 HTTP 伺服器

上網的時候，各位通常大概都發現過網站的網址最前面固定都會有一段 `http://` 或 `https://`。
我想，只有很少數的使用者思考過，這一段文字到底是什麼意思。

HTTP 是一種協議，也就是說，一種規定兩臺電腦之間的溝通方式的<a href="https://www.rfc-editor.org/rfc/rfc9110.html" target="_blank" rel="noopener noreferrer">文件</a>。
每一次想要透過 HTTP 協議得到一個資源，比如說當你打開此課程的時候，瀏覽器會在背後向某一臺伺服器「請求」每一個頁面。
對應每一個請求，伺服器將分別傳送一個「回應」。
提供資料的機器叫作伺服器（英：server），而請求資料的那一臺電腦則稱為客戶端（英：client）。

## 為 Git 設定開發者資料

請保證你的 Git 設定好了信箱與用戶名。以下是在我的電腦上執行 `git config --global --list` 的結果：

```shell
$ git config --global --list
core.excludesfile=/home/karol/.gitignore
user.email=karol@moroz.dev
user.name=Karol Moroz
init.defaultbranch=main
pull.rebase=false
```

如果所列印的結果中看到了 `user.email` 與 `user.name` 兩個設定值，代表設定正確。
如果這兩個設定沒出現，請用以下指令設定：

```shell
git config --global user.name "Xiaoming Wang" # 填寫自己的資料
git config --global user.email "xiaoming.wang@example.com"
```

以上資料將會儲存在每一個「提交」（英：commit）裡面。

## 建立課程資料夾

整個課程中，我們的練習題都會寫在同一個資料夾。
這個資料夾具體的路徑各位讀者當然可以自己選擇，但在課程的內文中，我都會假設它在 `~/projects/go-tutorial`（`~`是你的<a href="https://zh.wikipedia.org/zh-tw/%E5%AE%B6%E7%9B%AE%E5%BD%95" target="_blank" rel="noopener noreferrer">家目錄</a>）。

以下指令將建立一個資料夾：

```shell
mkdir -p ~/projects/go-tutorial
```

`mkdir` 為 _make directory_ 的縮寫，而 `-p`（`--parents`的縮寫）代表我們希望建立中間每一層的資料夾，也就是說，如果 `~/projects` 不存在，這個指令也就會順便建立 `~/projects`。

進入稍早建立的資料夾：

```shell
cd ~/projects/go-tutorial
```

接下來可以在此初始化一個 Git 版本庫：

```shell
git init
```

每一個 Git 版本庫都需要至少有一個「提交」，才可以上傳到 Github，而每一個提交都需要至少變更一個檔案。
所以在這個資料夾新增一個簡單的 `README.md`檔案，這是一個 Git 專案的介紹文件：

```markdown
# 網頁開發學習日記

大家好，我正在跟肉肉教授學習網頁開發！
請上 https://tutorial.moroz.dev 加入我！
```

如果在這個資料夾執行 `git status` 應該會看到類型以下的結果：

```shell
$ git status
On branch main

No commits yet

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        README.md

nothing added to commit but untracked files present (use "git add" to track)
```

將稍早建立的 `README.md` 加入第一個提交：

```shell
go add README.md
```

完成建立第一個提交：

```shell
$ git commit -m "Initial commit"
[main (root-commit) a72ac94] Initial commit
 1 file changed, 5 insertions(+)
 create mode 100644 README.md
```

現在可以至 Github 建立一個遠端版本庫並按照網頁上的指示上傳本地資料。

## 建立小專案

有了一個 Git 版本庫後，就可以建立第一個小程式專案。

請保證你在 `~/projects/go-tutorial` 資料夾中：

```shell
cd ~/projects/go-tutorial
```

然後建立名叫 `01-server` 的資料夾：

```shell
mkdir 01-server
```

進入稍早建立的資料夾：

```shell
cd 01-server
```

初始化新的 Go 專案：

```shell
go mod init github.com/<你的Github用戶名>/go-tutorial/01-server
```

## 第一個伺服器

在這個資料夾裡面，建立一個新檔案 `main.go` 並填寫以下內容：

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

這個函數的第一個參數為希望監聽的端口，而第二個參數為一個 `http.Handler`。
如果 `ListenAndServe` 發生錯誤，它將返回一個錯誤值（`error`），然後由 `log.Fatal` 印出來，小程式就結束。
如果沒有出錯，那麼 `http.ListenAndServe` 將持續監聽我們所指定的端口並使用 `handler` 處理 HTTP 請求。

### `http.Handler` 與 `http.HandlerFunc` 的用法

`http.ListenAndServe` 的第二個參數叫 `handler`，它的類型為 `http.Handler` 介面（英：interface）。
`http.Handler` 要求一個方法：

```shell
$ go doc http.Handler
package http // import "net/http"

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

`http` 這個 `package` 內另外提供了一種類型，名為 `http.HandlerFunc`：

```shell
$ go doc http.HandlerFunc
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
