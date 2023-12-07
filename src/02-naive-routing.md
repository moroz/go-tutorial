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
```

`mkdir -p` 將建立 `02-naive-routing` 這個資料夾。
由於加了 `-p` 這個設定（`--parents`），這個指令將建立每一個還不存在的 _parent directory_。
例如，如果 `02-naive-routing` 的上一層 `go-tutorial` 還不存在，它就會建立 `go-tutorial` 這個資料夾。
如果 `go-tutorial` 的上一層 `projects` 還不存在，它就會建立 `projects` 這個資料夾等等。

`-p` 這個設定也可以理解為「保證一個資料夾存在」。
沒有加 `-p` 的 `mkdir` 期待要建立的資料夾還不存在，如果它已經存在就會返回一個錯誤：

```shell
$ mkdir 02-naive-routing # 由於我剛剛已經建立了這個資料夾，它會出錯
mkdir: cannot create directory ‘02-naive-routing’: File exists
$ mkdir -p 02-naive-routing # 沒毛病
```

進入稍早建立的資料夾：

```shell
cd ~/projects/go-tutorial/02-naive-routing
```

初始化新的 Go 專案：

```shell
go mod init go-tutorial/02-naive-routing
```

然後在這個專案裡面新增 `main.go` 的檔案：

```go
{{#include ../code/02-naive-routing/iterations/01/main.go}}
```

這個小程式與上一堂課差不多一樣。首先，`main()` 的功能一樣：這個程式將監聽 3000 端口並用 `HandleRequest` 這個函數處理請求。
`HandleRequest` 目前是空的，我們要找辦法在 `HandleRequest` 裡針對不同路徑提供不同的內容。

## 尋找請求路徑

一個網頁的網址都符合一種統一的格式，名叫 URL（Uniform Resource Locator，統一資源定位符）。
各位同學肯定看過各式各樣的 URL，有的可能也聽過 URL 這個說法。
以下為一個簡單的 URL 範例：

```
https://tutorial.moroz.dev/02-naive-routing
```

這個 URL 可以進一步分析為：

```
https -- scheme 協定名，代表這個網頁用加密的 HTTPS 協議
tutorial.moroz.dev -- host 主機名，這個網站的域名
/02-naive-routing -- path 路徑，一律都是 / 開頭
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

賓果！我們找到了一個路徑！這代表我們可以根據 `r`（`*http.Request`）的 `URL`（`*url.URL`）的 `Path`（`string`）判斷請求路徑，提供不同內容：

```go
{{#include ../code/02-naive-routing/iterations/02/main.go:13:24}}
```

## `http.ResponseWriter` 的 `Write` 方法

在上一堂課裡面，我們的 `HandleRequest` 都只有顯示一樣的內容：

```go
{{#include ../code/01-server/main.go:13:15}}
```

從這一段程式可見，返回一個 HTML 回應的方式就是呼叫 `w` 的 `Write` 方法。
`w` 的類型為 `http.ResponseWriter`，我們來看看 `Write` 方法的文檔：

```go
$ go doc http.ResponseWriter.Write
package http // import "net/http"

type ResponseWriter interface {

	// Write writes the data to the connection as part of an HTTP reply.
	//
	// If WriteHeader has not yet been called, Write calls
	// WriteHeader(http.StatusOK) before writing the data. If the Header
	// does not contain a Content-Type line, Write adds a Content-Type set
	// to the result of passing the initial 512 bytes of written data to
	// DetectContentType. Additionally, if the total size of all written
	// data is under a few KB and there are no Flush calls, the
	// Content-Length header is added automatically.
	//
	// Depending on the HTTP protocol version and the client, calling
	// Write or WriteHeader may prevent future reads on the
	// Request.Body. For HTTP/1.x requests, handlers should read any
	// needed request body data before writing the response. Once the
	// headers have been flushed (due to either an explicit Flusher.Flush
	// call or writing enough data to trigger a flush), the request body
	// may be unavailable. For HTTP/2 requests, the Go HTTP server permits
	// handlers to continue to read the request body while concurrently
	// writing the response. However, such behavior may not be supported
	// by all HTTP/2 clients. Handlers should read before writing if
	// possible to maximize compatibility.
	Write([]byte) (int, error)
}
```

由此可見，如果在呼叫 `Write` 沒有使用 `WriteHeader`，`Write` 將會設定預設的回應狀態碼 `http.StatusOK`：

```go
$ go doc http.StatusOK | grep StatusOK
        StatusOK                   = 200 // RFC 9110, 15.3.1
```

以上代表，如果沒有指定其他回應狀態碼，Go 將會預設返回 200 OK，代表請求處理成功。
大多人都知道 404 代表「找不到頁面」，Go 有沒有相關常數？

```go
$ go doc http.StatusOK | grep 404
        StatusNotFound                     = 404 // RFC 9110, 15.5.5
```

HTTP 協議指定了幾十種狀態碼，要看 Go 標準函數庫所定義的狀態碼常數可以看 `go doc http.StatusOK`。
相關的標準文件為 <a href="https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml" target="_blank" rel="noopener noreferrer">Hypertext Transfer Protocol (HTTP) Status Code Registry</a>。
另外，想要進一步了解狀態碼的讀者可以看 <a href="https://developer.mozilla.org/zh-TW/docs/Web/HTTP/Status" target="_blank" rel="noopener noreferrer">MDN: HTTP 狀態碼</a>。

除了，`Write` 將會根據我們所寫入的內容前512個字元判斷所寫入的格式並設定回應的 <a href="https://developer.mozilla.org/zh-TW/docs/Web/HTTP/Headers/Content-Type" target="_blank" rel="noopener noreferrer">`Content-Type` 標頭</a>，這解釋了為什麼之前我們返回**看起來**像 HTML 的內容，瀏覽器就將它理解成 HTML 內容。
