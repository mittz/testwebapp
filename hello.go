package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(w, "Hello, World")
		fmt.Fprintf(w, hostname)
	} else {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":8080", nil)
}
