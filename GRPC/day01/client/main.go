package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello_grpc "grpccase/pb"
)

func main() {
	dial, err := grpc.Dial("127.0.0.1:6000", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	client := hello_grpc.NewUserInfoClient(dial)
	req, err := client.GetUserInfo(context.Background(), &hello_grpc.UserRequest{Name: "我从客户端来"})
	fmt.Println(req.GetName())
}
