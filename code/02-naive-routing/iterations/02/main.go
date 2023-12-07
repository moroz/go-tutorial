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
	case "/":
		// 顯示首頁

	case "/contact":
		// 顯示聯絡資訊網站

	default:
		// 忽略任何其他路徑：顯示 404 錯誤
	}
}
