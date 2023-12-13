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

這段程式碼是一非常簡單的使用 `chi` 的伺服器：

```go
{{#include ../code/03-chi-router/iterations/01/main.go}}
```

```go
{{#include ../code/03-chi-router/iterations/02/main.go}}
```

```go
{{#include ../code/03-chi-router/iterations/03/main.go}}
```


