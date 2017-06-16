package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	var message = "Good morning@" + hostname + "\n"
	fmt.Fprintf(w, message)
	var _ = hostname
	var _ = err
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
