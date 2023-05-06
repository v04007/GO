package main

import (
	"fmt"
	"net"
	"strings"
)

// server
func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 20000,
	})
	if err != nil {
		fmt.Println("listen UDP faile:", err)
		return
	}
	defer conn.Close() //先判断是否为空，若是为空就没有Close方法
	var data [1204]byte
	for {
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read from UDP failed:", err)
			return
		}
		fmt.Println(data[:n])
		msg := strings.ToUpper(string(data[:n]))
		conn.WriteToUDP([]byte(msg), addr)
	}
}
