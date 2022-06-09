package main

import (
	"context"
	"fmt"
	//"google.golang.org/grpc"
	"google.golang.org/grpc"
	hello_grpc "grpccase/pb"
	"net"
)

type server struct {
	hello_grpc.UnimplementedUserInfoServer
}

func (s *server) GetUserInfo(ctx context.Context, req *hello_grpc.UserRequest) (res *hello_grpc.UserResponse, err error) {
	fmt.Println(req.GetName())
	return &hello_grpc.UserResponse{Name: "我是从服务端返回的grpc的内容"}, nil
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	s := grpc.NewServer()
	hello_grpc.RegisterUserInfoServer(s, &server{})
	s.Serve(listen)
}
