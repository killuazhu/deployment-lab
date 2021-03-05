package main

import (
	"fmt"
	"net/http"
)

const version = "v1"

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello %v\n", version)
}

func health(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "up\n")
}

func main() {

	http.HandleFunc("/health", health)
	http.HandleFunc("/", hello)

	http.ListenAndServe(":8090", nil)
}
