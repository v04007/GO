package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是主页")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是用户")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是创建用户")
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "创建订单")
}

func main() {
	server := NewHttpServer("test-server")
	server.Router("get", "/user/signup", SignUp)
	//server.Router("/user", user)
	//server.Router("/user/create", createUser)
	//server.Router("/order", order)
	server.Start(":8080")
}
