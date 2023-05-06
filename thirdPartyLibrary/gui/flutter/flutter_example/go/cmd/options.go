package main

import (
	"github.com/bettersun/go-flutter-plugin/hello"
	//"flutter_example/go/go_plugin/hello"
	"github.com/go-flutter-desktop/go-flutter"
)

var options = []flutter.Option{
	flutter.WindowInitialDimensions(800, 1280),
	flutter.AddPlugin(hello.HelloPlugin{}),
}
