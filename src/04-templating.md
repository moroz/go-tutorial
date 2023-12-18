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

以上為一個簡單的 HTML 樣版，接下來我們來

```go
{{#include ../code/04-templating/iterations/01/main.go}}
```
