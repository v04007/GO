package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", header)
	http.HandleFunc("/healthz", healthz)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func header(w http.ResponseWriter, r *http.Request) {
	h := r.Header.Get("token")
	fmt.Fprintln(w, "token", h)
	w.WriteHeader(200)
	fmt.Println(r.Host)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	VERSION := os.Getenv("GOPATH")
	fmt.Fprintln(w, 200)
	fmt.Fprintln(w, VERSION)
}
