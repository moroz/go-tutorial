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
