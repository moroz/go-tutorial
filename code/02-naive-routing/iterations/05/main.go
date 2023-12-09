package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", http.HandlerFunc(HandleRequest)))
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	// 首頁
	case "/":
		w.Write([]byte(`
			<h1>歡迎光臨王小明的網站！</h1>
			<a href="/contact">聯絡</a>
		`))

	// 聯絡資訊網站
	case "/contact":
		w.Write([]byte(`
			<h1>聯絡資訊</h1>
			<a href="/">返回首頁</a>
		`))

	// 忽略任何其他路徑
	default:
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`
			<h1>404 Not Found</h1>
			<p>找不到此頁面</p>
			<a href="/">返回首頁</a>
		`))
	}
}
