package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.NotFound(notFoundHandler)

	log.Println("Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
		<h1>歡迎光臨王小明的網站！</h1>
		<a href="/contact">聯絡</a>
	`)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
		<h1>聯絡資訊</h1>
		<a href="/">返回首頁</a>
	`)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, `
		<h1>404 Not Found</h1>
		<p>找不到此頁面</p>
		<a href="/">返回首頁</a>
	`)
}
