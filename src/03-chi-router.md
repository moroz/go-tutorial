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

在 Go 專案裡安裝第三方軟體包的指令為 `go get`。
<!-- `go get` 的功能與 Git 版本控制系統息息相關，安裝軟體包最簡單的方式就是 -->
<a href="https://github.com/go-chi/chi" target="_blank" rel="noopener noreferrer">Chi 的官方網站</a>介紹了安裝該軟體包的指令：

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



```go
{{#include ../code/03-chi-router/iterations/03/main.go}}
```


