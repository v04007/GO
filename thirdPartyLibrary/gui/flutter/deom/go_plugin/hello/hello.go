package hello

import (
	"fmt"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

// go-flutter插件需要声明包名和函数名
// dart代码中调用时需要指定相应的包名和函数名
const (
	channelName = "instance.id/go/data"
	hello       = "hello"
)

type Example struct {
	channel *plugin.MethodChannel
}

// HelloPlugin 声明插件结构体
//type HelloPlugin struct{}

// 指定为go-flutter插件
//var _ flutter.Plugin = &HelloPlugin{}

// InitPlugin 初始化插件
func (p *Example) InitPlugin(messenger plugin.BinaryMessenger) error {
	p.channel = plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	p.channel.HandleFunc(hello, helloFunc)
	return nil
}

// 插件中具体的执行函数
func helloFunc(arguments interface{}) (reply interface{}, err error) {
	fmt.Println("hello 执行了")
	return fmt.Sprintf("hello go-flutter%s", arguments), nil
}

//package hello
//
//import (
//	"fmt"
//	"github.com/go-flutter-desktop/go-flutter"
//	"github.com/go-flutter-desktop/go-flutter/plugin"
//)
//
//// go-flutter插件需要声明包名和函数名
//// dart代码中调用时需要指定相应的包名和函数名
//const (
//	channelName = "bettersun.go-flutter.plugin.hello"
//	hello       = "hello"
//)
//
//// HelloPlugin 声明插件结构体
//type HelloPlugin struct{}
//
//// 指定为go-flutter插件
//var _ flutter.Plugin = &HelloPlugin{}
//
//// InitPlugin 初始化插件
//func (HelloPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
//	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
//	channel.HandleFunc(hello, helloFunc)
//	return nil
//}
//
//// 插件中具体的执行函数
//func helloFunc(arguments interface{}) (reply interface{}, err error) {
//	fmt.Println("hello 执行了")
//	return "hello go-flutter", nil
//}
