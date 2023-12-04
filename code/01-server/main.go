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
