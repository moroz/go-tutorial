# 第二章&#x3000;不同分頁

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

從使用者的角度來看，路徑就是每一個分頁的網址。
當你瀏覽到不同分頁的時候，瀏覽器將會幫你向伺服器請求每一個分頁的內容。
每一個分頁都用一個獨立的 HTTP 請求，請求資料包含路徑等資料。
伺服器將按照這路徑決定要回應什麼樣的資料給瀏覽器。

所以，開發不同分頁的時候，我們需要知道所請求的路徑，才可以知道要返回什麼資料。
這個專案裡面，負責路由（英：routing）的邏輯都寫在 `HandleRequest` 函數裡。
所以，寫這個函數的第一個步驟就是判斷請求路徑，然後用 `if` 或 `switch` 這種條件判斷表達式產生對應的內容。

`HandleRequest` 接受兩個參數，類型分別為 `http.ResponseWriter` 與 `*http.Request`（`http.Request` 的指標）。
我們需要的是請求（request）的資料。
`ResponseWriter` 這個類型的名稱為 response（回應）writer（寫入器），所以從名字上可以推理出，它與請求無關。
反而，`Request` 的名稱就是我們目前所需的「請求」，可以查詢看看它是否包含請求路徑。
這時候建議看相關文檔，在此為 `http.Request`。
以下指令將印出 `http.Request` 的所有屬性：

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

## 實作首頁與聯絡資訊

知道如何判斷請求路徑之後，我們就可以來寫出首頁與聯絡資訊頁面：

```go
{{#include ../code/02-naive-routing/iterations/03/main.go}}
```

現在，我們可以測試這個程式的功能：

```shell
$ go run .
2023/12/09 02:05:01 Listening on :3000...
```

<figure class="bordered-figure">
<a href="/images/02/main.webp" target="_blank" rel="noopener noreferrer"><img src="/images/02/main.webp" /></a>
<caption>執行以上程式，瀏覽至<code>http://localhost:3000</code>的畫面。</caption>
</figure>

<figure class="bordered-figure">
<a href="/images/02/contact.webp" target="_blank" rel="noopener noreferrer"><img src="/images/02/contact.webp" /></a>
<caption>點擊「聯絡」的連接後的畫面。</caption>
</figure>

然後，如果現在瀏覽到一個不存在的頁面的話，只看得到空白一片：

<figure class="bordered-figure">
<a href="/images/02/non-existent.webp" target="_blank" rel="noopener noreferrer"><img src="/images/02/non-existent.webp" /></a>
<caption>瀏覽至不存在的頁面後只看得到空白一片。</caption>
</figure>

大多數的讀者或許在網路上看過 404 這個數字，意思就是「找不到頁面」。
照理來說，如果使用者瀏覽到一個不存在的畫面的話，我們應該返回一個 404 錯誤頁面，並在頁面裡顯示首頁的連接。
然而，在我們能夠成功完成這個任務之前，應該要先明白錯誤狀態碼以及 `Write` 的使用方法。

## `http.ResponseWriter` 的 `Write` 方法

目前，我們已經實作的兩個畫面都用同樣的方式返回結果，也就是說，直接呼叫 `w` 這個變數的 `Write` 方法：

```go
{{#include ../code/02-naive-routing/iterations/03/main.go:15:20}}
```

我們知道 `HandleRequest` 接受兩個參數，而名叫 `w` 的那個參數類型為 `http.ResponseWriter`：

```go
{{#include ../code/02-naive-routing/iterations/03/main.go:13}}
```

以下指令可檢查 `http.ResponseWriter` 的 `Write` 方法的文檔：

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

### 狀態碼

`http.ResponseWriter` 的 `Write` 方法接受一個參數，其類型為 `[]byte`（二進制資料），並返回兩個值：`(int, error)`（一個整數與一個錯誤）。
目前，它的返回值對我們來說不重要，然而文檔裡第二段文字包含了重要的解釋：

> If `WriteHeader` has not yet been called, `Write` calls `WriteHeader(http.StatusOK)` before writing the data.

如果在呼叫 `Write` 之前沒有使用 `WriteHeader`，`Write` 將會呼叫該方法來設定回應狀態碼 `http.StatusOK`。
`http.StatusOK` 為 `http` 包所定義的常數。檢查 `http.StatusOK` 的文檔：

```
go doc http.StatusOK
```

由於 `http` 源代碼中，所有狀態碼常數都是一起定義的，因此執行以上指令也會同時印出 `http` 定義的所有狀態碼常數，內容有點太多，但可以用 `grep` 來搜尋：

```go
$ go doc http.StatusOK | grep StatusOK
        StatusOK                   = 200 // RFC 9110, 15.3.1
```

原來 `StatusOK` 代表狀態碼 200。
從 <a href="https://developer.mozilla.org/zh-TW/docs/Web/HTTP/Status/200" target="_blank" rel="noopener noreferrer">
MDN: 200 OK</a> 的文檔裡可以得知：

> HTTP `200 OK` 成功狀態碼表明請求成功。

如果 200 狀態碼代表 OK，那麼衆所周知的 404 代表什麼？

> HTTP `404 Not Found` 用戶端錯誤回應碼，表明了伺服器找不到請求的資源。
引發 404 頁面的連結，通常被稱作斷連或死連（broken or dead link）、並可以導到失效連結（link rot）頁面。
（摘錄自：<a href="https://developer.mozilla.org/zh-TW/docs/Web/HTTP/Status/404" target="_blank" rel="noopener noreferrer">MDN: 404 Not Found</a>）

至於 404 錯誤碼，Go 的 `http` 包有沒有定義相關常數？

```go
$ go doc http.StatusOK | grep 404
        StatusNotFound                     = 404 // RFC 9110, 15.5.5
```

HTTP 協議指定了幾十種狀態碼，想要進一步了解狀態碼的讀者可以看 <a href="https://developer.mozilla.org/zh-TW/docs/Web/HTTP/Status" target="_blank" rel="noopener noreferrer">MDN: HTTP 狀態碼</a>與相關標準 <a href="https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml" target="_blank" rel="noopener noreferrer">Hypertext Transfer Protocol (HTTP) Status Code Registry</a>。

### `Content-Type` 標頭

`http.ResponseWriter.Write` 除了設定預設狀態碼 200 OK 以外，還會猜測並設定一個 <a href="https://developer.mozilla.org/zh-TW/docs/Web/HTTP/Headers/Content-Type" target="_blank" rel="noopener noreferrer">`Content-Type` 標頭</a>：

> If the `Header` does not contain a `Content-Type` line, `Write` adds a `Content-Type` set to the result of passing the initial 512 bytes of written data to `DetectContentType`.

至於 `http.DetectContentType` 的功能：

```shell
$ go doc http.DetectContentType
package http // import "net/http"

func DetectContentType(data []byte) string
    DetectContentType implements the algorithm described at
    https://mimesniff.spec.whatwg.org/ to determine the Content-Type of
    the given data. It considers at most the first 512 bytes of data.
    DetectContentType always returns a valid MIME type: if it cannot determine a
    more specific one, it returns "application/octet-stream".

```

這個函數將檢查資料的前 512 字元，然後保證總是會返回一個正確的 <a href="https://developer.mozilla.org/zh-TW/docs/Glossary/MIME_type" target="_blank" rel="noopener noreferrer">MIME 類型</a>。
讓我們試試看它的猜測多準確。以下程式碼包含幾個字串，分別為 HTML、CSS、JavaScript 與 JSON 的範例（這四種語言與格式都是網頁開發者的家常便飯），測試看看 `http.DetectContentType` 能夠正確猜測出哪些格式：

```go
{{#include ../code/02-naive-routing/iterations/04/main.go}}
```

執行以上程式的結果：

```shell
$ go run .
HTML: text/html; charset=utf-8
CSS: text/plain; charset=utf-8
JS: text/plain; charset=utf-8
JSON: text/plain; charset=utf-8
```

這恐怕沒有很實用：只有 HTML 猜中了（`text/html`），其他都直接當成純文字（`text/plain`）。
另外，Go 還加上了 `; charset=utf-8` 一段，表明我們將傳送 <a href="https://zh.wikipedia.org/wiki/UTF-8" target="_blank" rel="noopener noreferrer">UTF-8</a> 編碼（<a href="https://zh.wikipedia.org/wiki/Unicode" target="_blank" rel="noopener noreferrer">Unicode</a>）的文字資料，讓我們可以正確傳送拉丁字母以外的符號，例如漢字、注音符號抑是表情圖案，避開亂碼的問題。

雖然 HTML 以外的格式都被 `http.DetectContentType` 誤解成純文字，但重點是 HTML 那一段被 Go 猜對了！
這就是為什麼在第一章，當我們返回一段**看起來**像 HTML 的內容，瀏覽器就將它理解成 HTML 內容。
其實，瀏覽器本身沒有在猜測任何事，而是 Go 的 `http.ResponseWriter.Write` 幫我們檢測並設定正確的 `Content-Type` 標頭。

## 404 路徑

我們現在知道了，要讓瀏覽器正確顯示 HTML 內容，就是要在回應裡設定正確的 `Content-Type` 標頭：`text/html; charset=utf-8`。
另外，我們知道了 404 這個數字是一種「回應狀態碼」，而用 Go 語言設定狀態碼的方法就是呼叫 `http.ResponseWriter.WriteHeader`。
這樣就可以試著寫 `switch` 條件判斷的 `default:` 一段：

```go
{{#include ../code/02-naive-routing/iterations/05/main.go:30:37}}
```

執行這個程式：

```
$ go run .
2023/12/09 19:26:57 Listening on :3000...
```

現在，首頁與聯絡資訊頁面都可以正常使用，而瀏覽至不存在的頁面時就會看到 404 分頁：

<figure class="bordered-figure">
<a href="/images/02/404-correct.webp" target="_blank" rel="noopener noreferrer"><img src="/images/02/404-correct.webp" /></a>
<caption>瀏覽到 <code>http://localhost:3000/test</code> 的畫面。</caption>
</figure>

以上截圖上看得到開發者工具的「網路」分頁，可以看見我們所開發的伺服器應對第一個請求（`GET /test`）回應了正確的狀態碼 `404 Not Found` 與正確的 `Content-Type` 標頭 `text/html; charset=utf-8`。

在 Chromium 基礎的瀏覽器裡（Google Chrome、MS Edge、Opera）與 Mozilla Firefox 兩種瀏覽器內，這個畫面都可以用 F12 打開，或是在網頁的任何地方按右鍵並選取「檢查」（Inspect）。
開發者工具預設開啓的分頁可能因瀏覽器類型而異，要檢查網路請求的記錄，請點擊「網路」（Network）分頁並刷新網頁（快捷鍵為 F5 或 Ctrl-R，蘋果為 Cmd-R）。

