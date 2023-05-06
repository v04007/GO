package hello

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

// go-flutter插件需要声明包名和函数名
// dart代码中调用时需要指定相应的包名和函数名
const (
	channelName = "bettersun.go-flutter.plugin.hello"
	hello       = "hello"
	message     = "message"
)

// HelloPlugin 声明插件结构体
type HelloPlugin struct{}

// 指定为go-flutter插件
var _ flutter.Plugin = &HelloPlugin{}

// InitPlugin 初始化插件
func (HelloPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc(hello, helloFunc)
	channel.HandleFunc(message, messageFunc)

	// 向Dart端发送消息
	go sendInterval(channel)

	return nil
}

// 具体的逻辑处理函数，无参数传递
func helloFunc(arguments interface{}) (reply interface{}, err error) {
	return "Hello, go-flutter", nil
}

// 具体的逻辑处理函数，有参数传递
func messageFunc(arguments interface{}) (reply interface{}, err error) {
	var param string
	if arguments == nil {
		param = ""
	}

	switch arguments.(type) {
	case string:
		param = arguments.(string)
	default:
		newValue, _ := json.Marshal(arguments)
		param = string(newValue)
	}

	return "Welcome to go-flutter, " + param, nil
}

var tcpch chan string
var connop chan string

// 向Dart端发送消息(定时)
func sendInterval(channel *plugin.MethodChannel) {
	var str, IP string
	tcpch = make(chan string)
	connop = make(chan string)
	m := make(map[interface{}]interface{})
	go tcpServer(tcpch, connop)

	go func() {
		for {
			time.Sleep(time.Second * 1)
			select {
			case str = <-tcpch:
				fmt.Println("chan 取值结果", str)
			case IP = <-connop:
				fmt.Println("IPchan 取值结果", IP)
			default:

			}
		}
	}()

	go func() {
		ticker := time.NewTicker(time.Second * 2)
		m[str] = str
		m[IP] = IP
		for _ = range ticker.C {
			err := channel.InvokeMethod("interval", m)
			if err != nil {
				log.Println(err)
			}
		}
	}()
}

func tcpServer(data, connop chan string) {
	const (
		ip   = "127.0.0.1"
		port = 4003
	)
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP(ip), Port: port})
	if err != nil {
		fmt.Println("监听端口失败:", err.Error())
		return
	}
	fmt.Println("已初始化连接，等待客户端连接...")

	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("接受客户端连接异常:", err.Error())
			continue
		}
		fmt.Println("客户端连接来自:", conn.RemoteAddr().String())
		connop <- conn.RemoteAddr().String()
		defer conn.Close()
		go handleConn(conn, connop)
	}
}

func handleConn(conn net.Conn, val chan string) {
	reader := bufio.NewReader(conn)
	for {
		time.Sleep(time.Second * 1)
		dat, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读取完毕。。。。。。")
			break
		}
		fmt.Println("读取信息=", dat)
		val <- dat
		fmt.Println("没有柱塞")
		conn.Write([]byte("end"))
	}
}
