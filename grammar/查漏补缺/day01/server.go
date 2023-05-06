package main

import "net/http"

//func main() {
//	hello := "你好"
//	num := utf8.RuneCountInString(hello)
//	fmt.Println(num)
//}

type Server interface {
	Routable
	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler *HandlerBasedOnMap
}

//闭包使用
//func (s *sdkHttpServer) Router(pattern string, handleFunc func(ctx *Context)) {
//	http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
//		ctx := NewContext(writer, request)
//		handleFunc(ctx)
//	})
//}

func (s *sdkHttpServer) Router(
	method string,
	pattern string,
	handleFunc func(ctx *Context)) {
	key := s.handler.key(method, pattern)
	s.handler.handlers[key] = handleFunc
}

func (s *sdkHttpServer) Start(address string) error {
	http.Handle("/", s.handler)
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}
