package main

import (
	"fmt"
	"net/http"
)

type Routable interface {
	Router(method string, pattern string, handleFunc func(ctx *Context))
}
type Handler interface {
	http.Handler
	Routable
}

type HandlerBasedOnMap struct {
	//key 应该是 method + url
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBasedOnMap) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := h.key(request.Method, request.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(NewContext(writer, request))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Not Found"))
	}
}

func (h *HandlerBasedOnMap) key(method, path string) string {
	return fmt.Sprintf("%s#%s", method, path)
}

var _Handler = &HandlerBasedOnMap{}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		handlers: make(map[string]func(ctx *Context)),
	}
}
