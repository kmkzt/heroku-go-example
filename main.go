package main

import (
				"io"
				"net/http"
				"os"
)

func root(w http.Response, r *http.Request) {
				io.WriteString(w, "Hello World!")
}

func main() {
				port: = os.GetEnv("PORT")
				http.HandleFunc("/", root)
				http.ListenAndServe("")
}
