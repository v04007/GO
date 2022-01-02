package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello_grpc "grpc01/pb"
)

func main()  {
	dial, err := grpc.Dial("localhost:8888",grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	client := hello_grpc.NewHelloGrpcClient(dial)
	req, err := client.Sayhi(context.Background(), &hello_grpc.Req{Message: "我从客户端来"})
	fmt.Println(req.GetMessage())
}