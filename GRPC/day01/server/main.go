package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello_grpc "grpc01/pb"
	"net"
)

type server struct {
	hello_grpc.UnimplementedHelloGrpcServer
}

func (s *server)Sayhi(ctx context.Context,req *hello_grpc.Req) (res *hello_grpc.Res,err error)  {
	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{Message: "我是从服务端返回的grpc的内容"},nil
}

func main()  {
	listen, err := net.Listen("tcp",":8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	s:=grpc.NewServer()
	hello_grpc.RegisterHelloGrpcServer(s,&server{})
	s.Serve(listen)
}