import "package:flutter/material.dart";
import 'package:robots_repository/robots_repository.dart';
import 'package:talos/pages/home/home_page.dart';
import 'package:talos/pages/home/pages/homestats_page.dart';
import 'package:webview_flutter/webview_flutter.dart';

import 'pages/homedetails_page.dart';
import 'pages/map_page.dart';
import 'pages/videostreaming_page.dart';

class DetailsScreen extends StatefulWidget {
  final RobotEntity robot;
  const DetailsScreen({super.key, required this.robot});

  @override
  State<DetailsScreen> createState() => _DetailsScreenState();
}

class _DetailsScreenState extends State<DetailsScreen> {
  int currentIndex = 1;
  late List<Widget> _widgetOptions;


  @override
  void initState() {
    _widgetOptions = <Widget>[
      VideoStreamingPage(),
      HomeDetailPage(robot: widget.robot),
      HomeStatsPage(robot: widget.robot),
      MapPage(robot: widget.robot)
    ];
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.lightBlue[50],
      appBar: AppBar(
        backgroundColor: Colors.lightBlue[50],
      ),
      body: _widgetOptions[currentIndex],
      
      bottomNavigationBar: BottomNavigationBar(
        type: BottomNavigationBarType.fixed,
        currentIndex: currentIndex,
        showUnselectedLabels: false,
        onTap:(index) => setState(() {
          currentIndex = index;
        }),
        items: [
          BottomNavigationBarItem(
            icon: Icon(Icons.video_call),
            label: 'VideoStreaming',
            backgroundColor: Colors.blue[500],
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.home),
            label: 'Home',
            backgroundColor: Colors.blue[500],
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.data_exploration),
            label: 'Data',
            backgroundColor: Colors.blue[500],
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.map),
            label: 'Map',
            backgroundColor: Colors.blue[500],
          ),
        ],
      ),
      //WebViewWidget(controller: controller)
    );
  }
}