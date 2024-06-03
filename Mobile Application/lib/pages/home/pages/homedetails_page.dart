import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:robots_repository/robots_repository.dart';
import 'package:talos/components/edit_item.dart';

class HomeDetailPage extends StatefulWidget {
  final RobotEntity robot;
  const HomeDetailPage({super.key, required this.robot});

  @override
  State<HomeDetailPage> createState() => _HomeDetailPageState();
}

class _HomeDetailPageState extends State<HomeDetailPage> {
  late TextEditingController _controllerName;
  late TextEditingController _controllerCreated;
  late TextEditingController _controllerUpdated;

  @override
  void initState(){
    super.initState();
    _controllerName = TextEditingController(text: widget.robot.robotname);
    _controllerCreated = TextEditingController(text: widget.robot.createdAt);
    _controllerUpdated = TextEditingController(text: widget.robot.updatedAt);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        backgroundColor: Colors.lightBlue[50],
        body: Padding(
          padding: const EdgeInsets.all(30),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const Text(
                "Statistics",
                style: TextStyle(
                  fontSize: 36,
                  fontWeight: FontWeight.bold,
                ),
              ),
              const SizedBox(height: 40),
              EditItem(
                title: "Name",
                widget: TextField(readOnly: true, controller: _controllerName),
              ),
              const SizedBox(height: 40),
              EditItem(
                title: "CreatedAt",
                widget: TextField(readOnly: true, controller: _controllerCreated),
              ),
              const SizedBox(height: 40),
              EditItem(
                title: "UpdatedAt",
                widget: TextField(readOnly: true, controller: _controllerUpdated),
              ),
            ],
          ),
        ));
  }
}
