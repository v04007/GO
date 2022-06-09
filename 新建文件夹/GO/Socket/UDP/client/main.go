package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// client
func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 20000,
	})
	if err != nil {
		fmt.Println("start server failed:", err)
		return
	}
	defer socket.Close()

	var reply [1204]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入内容：")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "eixt" {
			break
		}
		socket.Write([]byte(msg))
		//接受回复数据
		n, _, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("accept reply msg faile:", err)
			return
		}
		fmt.Println("接收到的信息：", string(reply[:n]))
	}

}
