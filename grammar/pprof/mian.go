package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

// https://www.liwenzhou.com/posts/Go/pprof/
// 访问http://localhost:8080/debug/pprof/ 查看
// go tool pprof profile 进入交互
func main() {
	http.HandleFunc("/pprof", pprofHandle)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func pprofHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pprof....")
}
