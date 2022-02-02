package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
	fmt.Fprintln(w, "Hello!")
}