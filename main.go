package main

import (
	"fmt"
	"net/http"
	"os"
)

const version = "v1"

var hostname = ""

func findHostname() {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	hostname = name
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%v: hello %v\n", hostname, version)
}

func health(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "up\n")
}

func main() {
	findHostname()
	http.HandleFunc("/health", health)
	http.HandleFunc("/", hello)

	http.ListenAndServe(":8090", nil)
}
