import 'dart:async';

import 'package:flutter/services.dart';
export 'hello_plugin.dart';

class HelloPlugin {
// go-flutter插件中的包名，两者必须一致
static const channel = const MethodChannel('instance.id/go/data');
// go-flutter插件中的函数名，无参
static Future<String> hello() =>  channel.invokeMethod("hello","传入的参数");
//
// // go-flutter插件中的函数名，有参
// static Future<String> message(String p) async => channel.invokeMethod("message", p);
}

