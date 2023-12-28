# 第四章&#x3000;`html/template` 樣版

在第三章裡，我們成功用 `chi-router` 開發了一個路由器，然後在瀏覽器成功顯示了一些 HTML 內容。
然而，目前為止內容都寫死在程式碼裡，這樣開發畫面未免不太方便。
因此，在這堂課我們來學習如何開發簡單的樣版（英：templates）。

使用樣版主要好處如下：

- 可以將內容切到不是程式碼的檔案，甚至在編譯後更改；
- 可以重複使用同樣的內容（像是頁腳或標頭）；
- 可以使用動態內容。

Go 語言標準庫提供了兩個處理樣版的包：`html/template` 與 `text/template`。
兩個包功能幾乎一樣，只是 `html/template` 多了一些確保 HTML 安全性的小功能。

## 初始化小專案

```shell
mkdir -p ~/projects/go-tutorial/04-templating
cd ~/projects/go-tutorial/04-templating
go mod init go-tutorial/04-templating
```

在 `04-templating` 資料夾時，建立一個新的資料夾 `templates`，我們將在裡面寫樣版檔案：

```shell
mkdir templates
```

在該新資料夾新增一個樣版檔案，檔案名為 `index.html.tmpl`：

```html
{{#include ../code/04-templating/iterations/01/templates/index.html.tmpl:1:5}}
```

以上為一個簡單的 HTML 樣版，接下來我們來將這個樣版顯示到終端機：

```go
{{#include ../code/04-templating/iterations/01/main.go}}
```

首先，`import` 的部分引用了 `html/template`。

```go
{{#include ../code/04-templating/iterations/01/main.go:3:6}}
```

使用 `import` 時應注意不要寫成 `text/template`，兩個包的函數都一樣，但渲染 HTML 的功能有所不同。

```go
{{#include ../code/04-templating/iterations/01/main.go:8:10}}
```

這邊定義了一個資料結構，該資料結構代表傳輸給樣版的資料。這邊的資料結構很簡單，只有一個叫 `Name` 的屬性，類型為字串。

在 `main` 函數裡，我們打開樣版的檔案並 parse 成樣版：

```go
{{#include ../code/04-templating/iterations/01/main.go:13:16}}
```

`template.ParseFiles` 返回兩個值：第一個值為 parse 好的樣版，而第二個為錯誤。
處理樣版對電腦來說是相當複雜的任務，第一次讀取一個樣版的時候需要驗證語法是否正確，將樣版轉換成可以快速執行的格式，有點像編譯。
如果這個過程中發生錯誤，例如指定的檔案不存在或是有語法錯，那麼第二個返回值就不是 `nil`，這種情況下接下來的程式無法正常執行，所以用 `panic` 終止程式。

```go
{{#include ../code/04-templating/iterations/01/main.go:17:19}}
```

用樣版的 `Execute` 方法將渲染結果印出到終端機。
Go 在這邊用 _execute_（執行）的說法，而不是 _render_（渲染）對某些人來說不太直覺。
以我的理解，它之所以叫 `Execute`，是因為在上一步（`ParseFiles`）我們將檔案裡的文字資料轉換成了可以快速執行的內部結構，渲染的時候真的有點像是執行一段程式。

`Execute` 的第一個參數為要輸出的地方，在此為終端機。
`os.Stdout` 指「作業系統」（英：**o**perating **s**ystem）的「標準輸出」（英：**st**andar**d out**put），在現代電腦系統中就是指終端機。

第二個參數為要傳輸給樣版的資料，它可以是任何類型，在此使用稍早專門定義的資料結構 `indexTemplateData`。

執行以上程式：

```shell
$ go run .
<title>Wang Xiaoming's Website</title>

<body>
  <h1>Hello, World</h1>
</body>
```

由於我們傳輸給樣版的資料裡有一個叫 `Name` 的屬性，它的值為字串「世界」，而樣版裡面用 `{{ .Name }}` 的寫法引用這個屬性，因此 `<h1>` 的內容變成「你好，世界」。

### 使用 `template.Must` 縮寫錯誤處理

在上方程式中，我們用這一段程式讀取樣版檔案，如果讀取過程中發生錯誤，將印出錯誤並提早終止程式：

```go
{{#include ../code/04-templating/iterations/01/main.go:13:16}}
```

一般而言，大部分的程式一定要所有樣版都讀取成功，才能夠正常使用，所以以上用法很常見。
`html/template` 提供了一個實用的小函數 `template.Must`：

```shell
$ go doc template.Must
package template // import "html/template"

func Must(t *Template, err error) *Template
    Must is a helper that wraps a call to a function returning (*Template,
    error) and panics if the error is non-nil. It is intended for use in
    variable initializations such as

        var t = template.Must(template.New("name").Parse("html"))
```

所以，我們程式可以寫成：

```go
{{#include ../code/04-templating/iterations/02/main.go:12:17}}
```

執行起來，結果還是一模一樣：

```shell
$ go run .
<title>Wang Xiaoming's Website</title>

<body>
  <h1>Hello, World</h1>
</body>
```

## 結合樣版與 HTTP 路由器

接下來，我們來開發一個簡單的網站，一樣有一個首頁與一個聯絡資訊頁面，還有客製化的 404 錯誤頁面。
這個小程式可以繼續用上一節的小專案用開發，但如果各位讀者想要保留第一個版本的程式碼，建議先 commit 到這的進度：

```shell
git add -A
git commit -m "Render template to STDOUT"
```

然後我們來修改現有的樣版。在 `templates/index.html.tmpl`，我們來新增聯絡資訊頁面的連接：

```html
<!-- templates/index.html.tmpl -->
{{#include ../code/04-templating/iterations/03/templates/index.html.tmpl}}
```

另外，在 `templates/contact.html.tmpl` 新增聯絡資訊頁面的樣版：

```html
<!-- templates/contact.html.tmpl -->
{{#include ../code/04-templating/iterations/03/templates/contact.html.tmpl}}
```

最後，在 `templates/404.html.tmpl` 新增 404 錯誤頁面的樣版：

```html
<!-- templates/404.html.tmpl -->
{{#include ../code/04-templating/iterations/03/templates/404.html.tmpl}}
```

```go
{{#include ../code/04-templating/iterations/03/main.go}}
```

<figure class="bordered-figure">
<a href="/images/04/main-template.webp" target="_blank" rel="noopener noreferrer"><img src="/images/04/main-template.webp" /></a>
<caption>使用<code>html/template</code>產生的首頁。</caption>
</figure>

<figure class="bordered-figure">
<a href="/images/04/contact-template.webp" target="_blank" rel="noopener noreferrer"><img src="/images/04/contact-template.webp" /></a>
<caption>使用<code>html/template</code>產生的聯絡資訊頁面。</caption>
</figure>

<figure class="bordered-figure">
<a href="/images/04/404-template.webp" target="_blank" rel="noopener noreferrer"><img src="/images/04/404-template.webp" /></a>
<caption>使用<code>html/template</code>產生的 404 錯誤頁面。</caption>
</figure>

