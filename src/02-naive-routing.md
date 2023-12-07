# 第二章&#x3000;請求與回應

目前，我們網頁進度還不錯：已經可以在瀏覽器裡面顯示內容了。
然而，這個網頁的功能還是有限：無論瀏覽到哪一個頁面，都只會顯示一樣的內容！
一般而言，網站都需要顯示多個畫面，例如一個首頁與一個聯絡資訊畫面。

各位同學可能發現過，一個網站的不同頁面通常會用不同的路徑。
因路徑而改變邏輯的功能稱為「路由」（英：routing），而負責路由的程式則稱為「路由器」（英：router）。
如果這個說法讓你想到家裡的 Wi-Fi 分享器，那就是因為這個機器的功能跟路由器的功能很像：網路分享器負責根據 IP 址將資料傳送到不同機器。

這堂課要開發的小程式將會顯示三個頁面：

* `GET /`：首頁。
* `GET /contact`：聯絡資訊頁面。
* 任何其他網址：404，找不到頁面。

讓我們建立一個新的小專案：

```shell
mkdir -p ~/projects/go-tutorial/02-naive-routing
cd ~/projects/go-tutorial/02-naive-routing
go mod init go-tutorial/02-naive-routing
```

```go
{{#include ../code/02-naive-routing/iterations/01/main.go}}
```
