package main

import (
	"github.com/bettersun/go-flutter-plugin/hello"
	"github.com/go-flutter-desktop/go-flutter"
)

var options = []flutter.Option{
	flutter.WindowInitialDimensions(800, 1280),
	// 添加插件
	flutter.AddPlugin(hello.HelloPlugin{}),
}
