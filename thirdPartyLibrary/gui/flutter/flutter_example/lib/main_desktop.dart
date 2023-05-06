import 'package:flutter/foundation.dart'
    show debugDefaultTargetPlatformOverride;
import 'package:flutter/material.dart';

void main() {
  debugDefaultTargetPlatformOverride = TargetPlatform.fuchsia;
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'go-flutter Texture API',
      theme: ThemeData(
        // If the host is missing some fonts, it can cause the
        // text to not be rendered or worse the app might crash.
        fontFamily: '这是什么',
        primarySwatch: Colors.blue,
      ),
      home: Scaffold(
        appBar: AppBar(
          title: Text('这是标题'), //title 标题
        ),
        body: Center(
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            children: <Widget>[
              Column(
                children: <Widget>[
                  Spacer(flex: 5),//距离上面位置
                  Text('图片上方标题'),
                  Spacer(flex: 2),
                  ConstrainedBox(
                    constraints: BoxConstraints.tight(Size(640, 560)), //图片高宽
                    // hard-coded to 2, gif.
                    // (Not the best practise, let go-flutter generate this ID
                    // and send it back to the dart code using platform
                    // messages不是最佳做法，让 go-flutter 生成此 ID 并使用平台消息将其发送回 dart 代码)
                    child: Texture(textureId: 2),
                  ),
                  Spacer(flex: 4),
                ],
              ),
              // Column(
              //   children: <Widget>[
              //     Spacer(flex: 2),
              //     Text('Image Texture (Cleared after 5s)'),
              //     Spacer(flex: 2),
              //     ConstrainedBox(
              //       constraints: BoxConstraints.tight(Size(330, 319)),
              //       child: Texture(textureId: 1), // hard-coded to 1, image硬编码为 1，图像
              //     ),
              //     Spacer(flex: 3),
              //   ],
              // ),//图片二
            ],
          ),
        ),
      ),
    );
  }
}
