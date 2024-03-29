import 'dart:async';

import 'package:flutter/services.dart';

class HelloPlugin {
// go-flutter插件中的包名，两者必须一致
static const channel =
MethodChannel("bettersun.go-flutter.plugin.hello");

// go-flutter插件中的函数名，无参
static Future<Future> hello() async => channel.invokeMethod("hello");

// go-flutter插件中的函数名，有参
static Future<Future> message(String p) async =>
channel.invokeMethod("message", p);
}