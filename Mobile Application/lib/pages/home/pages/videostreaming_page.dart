import 'package:flutter/material.dart';
import 'package:webview_flutter/webview_flutter.dart';

class VideoStreamingPage extends StatelessWidget {

  final controller = WebViewController()
  ..setJavaScriptMode(JavaScriptMode.unrestricted)
  ..setBackgroundColor(const Color(0x00000000))
  ..setNavigationDelegate(
    NavigationDelegate(
      onProgress: (int progress) {
        // Update loading bar.
      },
      onPageStarted: (String url) {},
      onPageFinished: (String url) {},
      onWebResourceError: (WebResourceError error) {},
      onNavigationRequest: (NavigationRequest request) {
        if (request.url.startsWith('https://www.youtube.com/')) {
          return NavigationDecision.prevent;
        }
        return NavigationDecision.navigate;
      },
    ),
  )
  //TODO: Robot url in robot object
  ..loadRequest(Uri.parse('http://nattech.fib.upc.edu:40344/video_feed/video'));

  @override
  Widget build(BuildContext context){
    return Scaffold(
      body: WebViewWidget(controller: controller),
    );
  }
}