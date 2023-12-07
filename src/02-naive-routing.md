# 第二章&#x3000;請求與回應

目前，我們網頁進度還不錯：已經可以在瀏覽器裡面顯示內容了。
然而，這個網頁的功能還是有限：無論瀏覽到哪一個頁面，都只會顯示一樣的內容！
一般而言，網站都需要顯示多個畫面，例如一個首頁與一個聯絡資訊畫面。

各位同學可能發現過，一個網站的不同頁面通常會用不同的路徑。
為不同路徑提供不同內容的功能稱為「路由」（英：routing），而負責路由的程式則稱為「路由器」（英：router）。
如果這個說法讓你想到家裡的 Wi-Fi 分享器，那就是因為這個機器的功能跟路由器的功能很像：網路分享器負責根據 IP 址將資料傳送到不同機器。

## 專案介紹

這堂課要開發的小程式將會顯示三個頁面：

* `/`：首頁。
* `/contact`：聯絡資訊頁面。
* 任何其他網址：找不到頁面，又名 404。

讓我們建立一個新的小專案：

```shell
mkdir -p ~/projects/go-tutorial/02-naive-routing
cd ~/projects/go-tutorial/02-naive-routing
go mod init go-tutorial/02-naive-routing
```

然後在這個專案裡面新增 `main.go` 的檔案：

```go
{{#include ../code/02-naive-routing/iterations/01/main.go}}
```

這個小程式與上一堂課差不多一樣。首先，`main()` 的功能一樣：這個程式將監聽 3000 端口並用 `HandleRequest` 這個函數處理請求。
`HandleRequest` 目前是空的，我們要找辦法在 `HandleRequest` 裡針對不同路徑提供不同的內容。

## 查詢請求路徑

一個網頁的網址都符合一種統一的格式，名叫 URL（Uniform Resource Locator，統一資源定位符）。
各位同學肯定看過各式各樣的 URL，有的可能也聽過 URL 這個說法。
以下為一個簡單的 URL 範例：

```
https://tutorial.moroz.dev/02-naive-routing
```

這個 URL 可以進一步分析為：

```
https -- scheme（協定名）
tutorial.moroz.dev -- host（主機名）
/02-naive-routing -- path（路徑）
```

我們在 `HandleRequest` 裡面將會需要請求路徑（request path），一般而言在文檔中搜尋英文的關鍵字是還不錯的查詢資料的策略。
`HandleRequest` 接受兩個參數，類型分別為 `http.ResponseWriter` 與 `*http.Request`（`http.Request` 的指標）。
`ResponseWriter` 的名稱為 response（回應）writer（寫入器），看起來與我們要找的資料無關。
反而，`Request` 就是我們需要的「請求」，我們可以看看有沒有提供路徑資訊。
各位同學可以自己執行以下指令來看看 `http.Request` 的屬性：

```shell
$ go doc http.Request
```

可見這個資料結構屬性很多，如果直接用看的會有點慢，但我們可以用另外一個指令來篩選此指令的輸出資料：

```shell
# grep 可以拿來在 go doc 裡搜尋 path 這個關鍵字
# -i (case-insensitive) 不分大小寫的搜尋
# -C10 (context) 顯示關鍵字上下各十行內容

$ go doc http.Request | grep -C10 -i path
        // Go's HTTP client does not support sending a request with
        // the CONNECT method. See the documentation on Transport for
        // details.
        Method string

        // URL specifies either the URI being requested (for server
        // requests) or the URL to access (for client requests).
        //
        // For server requests, the URL is parsed from the URI
        // supplied on the Request-Line as stored in RequestURI.  For
        // most requests, fields other than Path and RawQuery will be
        // empty. (See RFC 7230, Section 5.3)
        //
        // For client requests, the URL's Host specifies the server to
        // connect to, while the Request's Host field optionally
        // specifies the Host header value to send in the HTTP
        // request.
        URL *url.URL

        // The protocol version for incoming server requests.
        //
```

上方搜尋結果可見，`http.Request` 本身沒有叫 `Path` 的屬性，但是有叫 `URL` 的屬性，其類型為 `*url.URL`（`url.URL` 的指標）。
以下指令可以看 `url.URL` 的 `Path` 屬性的定義：

```go
$ go doc url.URL.Path             
package url // import "net/url"

type URL struct {
    Path string  // path (relative paths may omit leading slash)

    // ... other fields elided ...
}
```

賓果！我們找到了一個路徑！這代表我們可以根據 `r`（`*http.Request`）的 `URL`（`*url.URL`）的 `Path`（`string`）判斷請求路徑！

至於 `HandleRequest` 呢，差別在於，上一個程式對每一個請求都返回一模一樣的回應，而這個程式使用 `switch` 條件判斷，根據 `r.URL.Path` 的值改變邏輯：

```go
{{#include ../code/02-naive-routing/iterations/02/main.go:13:24}}
```

何謂 `r.URL.Path`？`r` 為 `HandleRequest` 的第二個參數，它的類型為 `http.Request`，它是一種描述 HTTP 請求的資料結構（`struct`）。
我們可以使用 `go doc` 來檢查 `http.Request` 的屬性：

```go
$ go doc http.Request.URL                                                        1
package http // import "net/http"

type Request struct {
    // URL specifies either the URI being requested (for server requests) or the URL to
    // access (for client requests).
    // 
    // For server requests, the URL is parsed from the URI supplied on the Request-Line
    // as stored in RequestURI. For most requests, fields other than Path and RawQuery
    // will be empty. (See RFC 7230, Section 5.3)
    // 
    // For client requests, the URL's Host specifies the server to connect to, while
    // the Request's Host field optionally specifies the Host header value to send in
    // the HTTP request.
    URL *url.URL

    // ... other fields elided ...
}
```

雖然這個英文描寫有一點模糊，但重點在這一段：

> URL specifies either the URI being requested (for server requests) or the URL to
> access (for client requests).

<!-- URI 為 Uniform Resource Identifier（統一資源標識符）的縮寫，而一個網頁的網址是 URL， -->
URL 就是一個網頁的網址，就像各位正在閱讀的頁面：

```
https://tutorial.moroz.dev/02-naive-routing
```

這個 URL 可以分析得更細：

```
https -- scheme（協定名）
tutorial.moroz.dev -- host（主機名）
/02-naive-routing -- path（路徑）
```

在這個小程式裡面，為了進行路由，只需要請求的路徑（path）。
`HandleRequest` 的第二個參數 `r` 類型為 `*http.Request`（`http.Request` 的指標），因此我們查得出來它有一個名叫 `URL` 的屬性，而這個屬性的類型
`http.Request` 的 `.URL` 屬性為一個 `*url.URL` 結構的指標，我們可以
