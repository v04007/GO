package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// TCP client

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("Dial failed:", err)
		return
	}
	Reader := bufio.NewReader(os.Stdin)
	var tmp [128]byte
	for {
		fmt.Println("请发送：")
		msg, _ := Reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			return
		}
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("write failed:", err)
			break
		}

		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from client failed,err:", err)
			return
		}
		fmt.Println("接收数据：", string(tmp[:n]))

	}
}
