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
<title>王小明的網站</title>

<body>
  <h1>你好，世界</h1>
</body>
```
