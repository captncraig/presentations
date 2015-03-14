package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, 世界!"))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
