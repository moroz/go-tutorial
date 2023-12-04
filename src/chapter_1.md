# 第一章 簡單的 HTTP 伺服器

上網的時候，各位通常大概都發現過網站的網址最前面固定都會有一段 `http://` 或 `https://`。
我想，只有很少數的使用者思考過，這一段文字到底是什麼意思。

HTTP 是一種協議，也就是說，一種規定兩臺電腦之間的溝通方式的<a href="https://www.rfc-editor.org/rfc/rfc9110.html" target="_blank" rel="noopener noreferrer">文件</a>。
每一次想要透過 HTTP 協議得到一個資源，比如說當你打開此課程的時候，瀏覽器會在背後向某一臺伺服器「請求」每一個頁面。
對應每一個請求，伺服器將分別傳送一個「回應」。
提供資料的機器叫作伺服器（英：server），而請求資料的那一臺電腦則稱為客戶端（英：client）。

尚未接觸 HTTP 協議的同學可能會對

## 初始化專案

```shell
mkdir -p ~/projects/go-tutorial/01-server
cd ~/projects/go-tutorial/01-server
```

```shell
go mod init go-tutorial/01-server
```

```go
package main

import (
	"log"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello from HandleRequest!</h1>"))
}

func main() {
	log.Println("Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", http.HandlerFunc(HandleRequest)))
}
```

<figure>
<img src="/images/01/01.png" />
<caption>執行以上程式，瀏覽至<code>http://localhost:3000</code>的畫面。</caption>
</figure>
