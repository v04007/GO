import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:hellogoflutter/plugin/go/plugin.dart';

class HelloPage extends StatefulWidget {
  @override
  _HelloPageState createState() => _HelloPageState();
}

class _HelloPageState extends State<HelloPage> {
  String timeNow = '';
  String IP = '';
  @override
  void initState() {
    // 监听Go端发过来的消息
    HelloPlugin.channel.setMethodCallHandler((MethodCall methodCall) async {
      if (methodCall.method == 'interval') {
        setState(() {
          print("methodCall===");
          print(methodCall.arguments);
          timeNow = methodCall.arguments;
        });
      }
    });

    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Row(
          children: [
            Text('连接设备:'),
            Container(width: 10),
            Text(IP),
          ],
        ),
      ),
      body: Center(
          child: Column(
        children: <Widget>[
          FutureBuilder<String>(
            future: HelloPlugin.hello(),
            builder: (c, snapshot) {
              if (!snapshot.hasData) {
                return Text('Hello插件hello函数运行出错');
              }
              return Text(timeNow);
            },
          ),
          FutureBuilder<String>(
            future: HelloPlugin.message('warrior'),
            builder: (c, snapshot) {
              if (!snapshot.hasData) {
                return Text('Hello插件message函数运行出错');
              }
              return Text(snapshot.data);
            },
          ),
        ],
      )),
    );
  }
}
