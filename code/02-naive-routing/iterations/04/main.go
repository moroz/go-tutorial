package main

import (
	"fmt"
	"net/http"
)

// HTML 的範例：這就是我們的首頁
const SAMPLE_HTML = `
<h1>Welcome to Wang Xiaoming's Website!</h1>
<a href="/contact">Contact</a>
`

// 簡單的 CSS 樣式範例
const SAMPLE_CSS = `
body {
	background: salmon;
}
`

// 簡單的 JavaScript
const SAMPLE_JS = `
import React from "react";
console.log("Hello, world!");
`

// JSON 資料
const SAMPLE_JSON = `{"hello":"World!"}`

func main() {
	fmt.Println("HTML:", http.DetectContentType([]byte(SAMPLE_HTML)))
	fmt.Println("CSS:", http.DetectContentType([]byte(SAMPLE_CSS)))
	fmt.Println("JS:", http.DetectContentType([]byte(SAMPLE_JS)))
	fmt.Println("JSON:", http.DetectContentType([]byte(SAMPLE_JSON)))
}
