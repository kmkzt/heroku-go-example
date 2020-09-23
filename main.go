package main

import (
	"io"
	"net/http"
	"os"
)

func root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func main() {
	// port := "3000"
	port := os.Getenv("PORT")
	http.HandleFunc("/", root)
	http.ListenAndServe(":"+port, nil)
}
