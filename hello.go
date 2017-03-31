package main

import (
	"fmt"
	"net/http"
	"log"
	"os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := exec.Command("hostname").Output()
	if err != nil {
		fmt.Fprintf(w, "Hello, World %s", hostname)
	} else {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":8080", nil)
}
