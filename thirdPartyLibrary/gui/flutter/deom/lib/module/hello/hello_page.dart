import 'package:flutter/material.dart';
import 'package:deom/plugin/go/hello_plugin.dart';

class HelloPage extends StatefulWidget {
  @override
  _HelloPageState createState() => _HelloPageState();
}

class _HelloPageState extends State<HelloPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Hello'),
      ),
      body: Center(
          child:  Column(
            children: <Widget>[
              FutureBuilder<String>(
                future: HelloPlugin.hello,
                builder: (c, snapshot) {
                  print("c===================");
                  print(c);
                  print("snapshot===================");
                  print(snapshot);
                  print("snapshot.hasData===");
                  print(snapshot.hasData);
                  if (!snapshot.hasData) {
                    return const Text('Hello插件执行出错');
                  }
                  return Text("snapshot.data");
                },
              )
            ],
          )),
    );
  }
}