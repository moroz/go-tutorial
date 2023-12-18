# 第三章&#x3000;Chi router

在第二章裡，我們實用 `switch` 條件判斷表達式開發了一款具有路由功能的小伺服器。
對於非常簡單的應用程式而言，使用一個 `switch` 可能是一個可行的作法，然而真正的網頁應用程式通常都有幾百，甚至可能是幾萬個獨立的頁面。
不僅如此，在一般的網頁應用裡並不是每一個人都可以直接拜訪每一個頁面，而會需要使用信箱與密碼登入，或是有更進階的權限管理，像是：「只有管理員才看得到這個頁面」。
當路由規則變得如此複雜的時候，對應的路由邏輯當然也就會變得非常複雜，如果一直使用單純的 `switch` 寫路由器，程式就會變得很冗長，不好維護。

所以，在這堂課我希望給各位讀者展示如何使用一個名叫 <a href="https://go-chi.io/#/" target="_blank" rel="noopener noreferrer">chi</a> 的第三方的軟體包。
Chi 是一種路由器，使用該軟體包就可以將路由器的程式拆成很多獨立的規則，甚至拆成多個獨立的小路由器，負責獨立的一個功能。

## 初始化小專案

使用以下終端機指令建立資料夾並初始化新專案。`mkdir -p` 的功能詳見第一章與第二章。
`go mod init` 將在當前資料夾初始化新的 Go 專案，由於我們在這堂課要使用第三方的軟體包，這個專案一定要執行 `go mod init` 才能正確安裝並編譯第三方軟體包。

```shell
mkdir -p ~/projects/go-tutorial/03-chi-router
cd ~/projects/go-tutorial/03-chi-router
go mod init go-tutorial/03-chi-router
```

在 Go 專案裡安裝第三方軟體包的指令為 `go get`。<a href="https://github.com/go-chi/chi" target="_blank" rel="noopener noreferrer">Chi 的官方網站</a>介紹了安裝該軟體包的指令：

```shell
go get -u github.com/go-chi/chi/v5
```

在稍早建立的專案資料夾中執行以上指令：

```shell
$ go get -u github.com/go-chi/chi/v5

go: downloading github.com/go-chi/chi v1.5.5
go: added github.com/go-chi/chi/v5 v5.0.10
```

這時候可能會發現 `go.mod` 檔案裡多了一行資料：

```plain
module go-tutorial/03-chi-router

go 1.21.1

require github.com/go-chi/chi/v5 v5.0.10 // indirect
```

每次在一個專案裡新增相依關係時，Go 將相依關係的資料寫入 `go.mod` 裡，保證之後在任何其他電腦上都可以正確安裝。

## `chi` 的用法

<a href="TODO" target="_blank" rel="noopener noreferrer">Chi 的網站</a>提供了還不錯的使用範例。
簡單地來講，`chi.NewRouter` 可以初始化一個路由器：

```go
r := chi.NewRouter()
```

以下指令可以查詢 `chi.NewRouter` 的文檔：

```shell
$ go doc github.com/go-chi/chi/v5 NewRouter
package chi // import "github.com/go-chi/chi/v5"

func NewRouter() *Mux
    NewRouter returns a new Mux object that implements the Router interface.
```

`chi.NewRouter` 反回一個 `*chi.Mux`（`chi.Mux` 結構的指標）。`chi.Mux` 提供了 `ServeHTTP(w http.ResponseWriter, r *http.Request)` 的方法，代表它符合 `http.Handler` 介面，可以直接用在 `http.ListenAndServe` 當 `handler` 參數：

```shell
$ go doc github.com/go-chi/chi/v5 Mux.ServeHTTP
package chi // import "github.com/go-chi/chi/v5"

func (mx *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request)
    ServeHTTP is the single method of the http.Handler interface that makes
    Mux interoperable with the standard library. It uses a sync.Pool to get and
    reuse routing contexts for each request.
```

以下程式將利用簡單的 `chi.Mux` 處理請求：

```go
{{#include ../code/03-chi-router/iterations/01/main.go}}
```

執行以上程式碼就會看到熟悉的畫面：

```shell
$ go run .
2023/12/16 17:06:22 Listening on :3000...
```

由於還沒有新增任何路徑，瀏覽到任何網址徑都是 404：

<figure class="bordered-figure">
<a href="/images/03/chi-no-routes.png" target="_blank" rel="noopener noreferrer"><img src="/images/03/chi-no-routes.png" /></a>
<caption>執行以上程式，瀏覽至<code>http://localhost:3000</code>的任何路徑，都只看得到 404 錯誤頁面。</caption>
</figure>

以下程式碼可以新增首頁與聯絡資訊兩個路徑：

```go
{{#include ../code/03-chi-router/iterations/02/main.go}}
```

以下程式碼建立新的 `chi.Mux` 並啟動請求記錄：

```go
{{#include ../code/03-chi-router/iterations/02/main.go:13:14}}
```

`middleware.Logger` 這一段的目的是為了加上記錄請求的功能。
它會讓我們的程式在終端機裡印出每一個請求的來源、路徑與回應狀態碼，如下：

```shell
$ go run .
2023/12/16 22:04:25 Listening on :3000...
2023/12/16 22:04:39 "GET http://localhost:3000/ HTTP/1.1" from 127.0.0.1:53156 - 200 79B in 44.334µs
```

以下兩行將新增 `/` 與 `/contact` 兩個路徑：

```go
{{#include ../code/03-chi-router/iterations/02/main.go:16:17}}
```

為什麼這兩個路徑都用 `Get` 方法定義呢？
以下為 `Mux.Get` 的文檔：

```shell
$ go doc github.com/go-chi/chi/v5 Mux.Get
package chi // import "github.com/go-chi/chi/v5"

func (mx *Mux) Get(pattern string, handlerFn http.HandlerFunc)
    Get adds the route `pattern` that matches a GET http method to execute the
    `handlerFn` http.HandlerFunc.
```

這個方法為什麼叫 `Get` 呢？
每一個 HTTP 請求除了路徑以外，都還有一個 method，最基本的兩個叫 `GET` 與 `POST`。
一般而言，當你使用瀏覽器拜訪一個頁面，或是點一個連接，瀏覽器都會替你傳送一個 `GET` 請求。
`POST` 請求主要用於送出表單，如寄信、登入等功能。
由於這兩個路徑都是用 `Get` 方法定義的，所以這兩個路徑都限使用 `GET` 這個 method，如果使用 `POST`，路由器將會返回 404 回應。

## Handler 函數

`Get` 的兩個參數分別為 `pattern` 與 `handlerFn`。
`pattern` 就是要處理的路徑，而 `handlerFn` 為處理請求的 `http.HandlerFunc`。
與 `http.ListenAndServe` 不同，這邊可以直接用函數，不需要專用 `http.Handler` 介面。
以下就是處理路徑的兩個函數：

```go
{{#include ../code/03-chi-router/iterations/02/main.go:23:35}}
```

這兩個函數都是我們已經熟悉的 `http.HandlerFunc`，因此兩個都接受同樣的參數，`(w http.ResponseWriter, r *http.Request)`。
但與上一堂課不同，我們用了一個協助寫資料的函數，`fmt.Fprint`。
以下為 `fmt.Fprint` 的文檔：

```shell
$ go doc fmt.Fprint
package fmt // import "fmt"

func Fprint(w io.Writer, a ...any) (n int, err error)
    Fprint formats using the default formats for its operands and writes to w.
    Spaces are added between operands when neither is a string. It returns the
    number of bytes written and any write error encountered.
```

`fmt.Fprint` 函數可以寫入任何一個符合 `io.Writer` 介面的物件。
`io.Writer` 的概念很簡單：不論輸出裝置是什麼，寫入資料的動作都一樣。
在 Go 語言裡面，可以寫入資料的物件都實作同樣的 `Write([]byte)` 的方法，這也是 `io.Writer` 介面要求唯一的方法。
因此，任何實作 `Write([]byte)` 方法的物件都可以當作 `io.Writer` 物件。

`homeHandler` 與 `contactHandler` 的第一個參數為一個 `http.ResponseWriter`。
我們知道 `http.ResponseWriter` 實作 `Write([]byte)` 的方法，這代表 `http.ResponseWriter` 也可以當作 `io.Writer`。
使用 `fmt.Fprint` 的好處是不再需要將回應內容轉換為 `string`。

## 預設路徑

目前我們的程式功能已經很接近第二章所開發的小專案，唯一的差別是還沒有設計客製化的 404 錯誤頁面。
如果瀏覽到不存在的路徑，就會看到預設 404 錯誤頁面，如圖：

<figure class="bordered-figure">
<a href="/images/03/chi-404.png" target="_blank" rel="noopener noreferrer"><img src="/images/03/chi-404.png" /></a>
<caption>瀏覽至不存在的路徑，就會看到預設的 404 錯誤頁面。</caption>
</figure>

為處理未知路徑的請求，`chi.Mux` 提供了一個方法，名叫 `NotFound`：

```shell
$ go doc github.com/go-chi/chi/v5 Mux.NotFound
package chi // import "github.com/go-chi/chi/v5"

func (mx *Mux) NotFound(handlerFn http.HandlerFunc)
    NotFound sets a custom http.HandlerFunc for routing paths that could not be
    found. The default 404 handler is `http.NotFound`.
```

在 `main()` 裡新增 `r.NotFound(notFoundHandler)` 一行：

```go
{{#include ../code/03-chi-router/iterations/03/main.go:12:21}}
```

然後新增 `notFoundHandler` 函數的定義：

```go
{{#include ../code/03-chi-router/iterations/03/main.go:37:}}
```

執行這段程式：

```shell
$ go run .
```

<figure class="bordered-figure">
<a href="/images/03/chi-custom-404.png" target="_blank" rel="noopener noreferrer"><img src="/images/03/chi-custom-404.png" /></a>
<caption>瀏覽至不存在的路徑，就會看到與第二章小專案一樣的客製化 404 錯誤頁面。</caption>
</figure>

這堂課的小專案就到這，我們來將更改儲存至 Git 版本庫：

```shell
git add .
git commit -m "03: Routing with chi-router"
```
