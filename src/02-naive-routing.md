# 第二章&#x3000;請求與回應

目前，我們網頁進度還不錯：已經可以顯示一些內容。
然而，它的功能還是有限：無論瀏覽到哪一個頁面，都顯示一樣的內容！
一般而言，網站都需要顯示多個畫面，例如一個首頁與一個聯絡資訊畫面。

讓我們建立一個新的小專案：

```shell
mkdir -p ~/projects/go-tutorial/02-naive-routing
cd ~/projects/go-tutorial/02-naive-routing
go mod init go-tutorial/02-naive-routing
```

```go
{{#include ../code/02-naive-routing/iterations/01/main.go}}
```
