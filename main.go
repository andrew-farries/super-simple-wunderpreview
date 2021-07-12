package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	hostname, _ := os.Hostname()

	fmt.Fprintf(w, "hello from: %s\n", hostname)
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Listening on port 8000")
	http.ListenAndServe(":8000", nil)
}
