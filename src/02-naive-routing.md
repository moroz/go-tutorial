# 第二章&#x3000;請求與回應

目前，我們網頁進度還不錯：已經可以在瀏覽器裡面顯示內容了。
然而，這個網頁的功能還是有限：無論瀏覽到哪一個頁面，都只會顯示一樣的內容！
一般而言，網站都需要顯示多個畫面，例如一個首頁與一個聯絡資訊畫面。

各位同學可能發現過，一個網站的不同頁面通常會用不同的路徑。
為不同路徑提供不同內容的功能稱為「路由」（英：routing），而負責路由的程式則稱為「路由器」（英：router）。
如果這個說法讓你想到家裡的 Wi-Fi 分享器，那就是因為這個機器的功能跟路由器的功能很像：網路分享器負責根據 IP 址將資料傳送到不同機器。

這堂課要開發的小程式將會顯示三個頁面：

* `/`：首頁。
* `/contact`：聯絡資訊頁面。
* 任何其他網址：找不到頁面，又名 404。

首先要明白，如何才能查詢一個請求的路徑。
一個網頁的網址都有一種統一的格式，該格式叫作 URL（Uniform Resource Locator，統一資源定位符）。
各位同學肯定看過各式各樣的 URL，範例如下：

```
https://tutorial.moroz.dev/02-naive-routing
```

這個 URL 可以進一步分析為：

```
https -- scheme（協定名）
tutorial.moroz.dev -- host（主機名）
/02-naive-routing -- path（路徑）
```

<!-- 至於404這個數字，我想大多數同學都熟悉這個錯誤號碼。 -->
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
至於 `HandleRequest` 呢，差別在於，上一個程式對每一個請求都返回一模一樣的回應，而這個程式使用 `switch` 條件判斷，根據 `r.URL.Path` 的值改變邏輯：

```go
{{#include ../code/02-naive-routing/iterations/01/main.go:13:17}}

    // ... 處理其他路徑
{{#include ../code/02-naive-routing/iterations/01/main.go:26:27}}
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
