package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// TCP server

func process(conn net.Conn) {
	var tmp [128]byte
	Reader := bufio.NewReader(os.Stdin)
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("conn read failed:", err)
			return
		}
		fmt.Println("接收数据：", string(tmp[:n]))
		fmt.Println("请回复：")
		msg, _ := Reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start listener 127.0.0.1:20000 failed :", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed:", err)
			return
		}
		go process(conn)
	}
}
