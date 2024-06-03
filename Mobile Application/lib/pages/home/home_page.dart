import 'dart:developer';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:talos/pages/home/blocs/get_robots_bloc/get_robots_bloc.dart';
import 'package:talos/pages/home/robot_details_page.dart';

class HomePage extends StatelessWidget {
  const HomePage({super.key});

  void signUserOut() async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.remove('token');
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.lightBlue[50],
      appBar: AppBar(
        backgroundColor: Colors.lightBlue[50],
        actions: [IconButton(onPressed: signUserOut, icon: const Icon(Icons.logout))],
        title: const Row(
          children: [
            //TODO: Posar imatge petita //Image.asset('lib/images/logo.png', scale: 1),
            Text(
              'Dashboard',
              style: TextStyle(fontWeight: FontWeight.w900, fontSize: 30)
            )
          ],
        ),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: BlocBuilder<GetRobotsBloc, GetRobotsState>(
          builder: (context, state) {
            if(state is GetRobotSuccess){
            return RefreshIndicator(
              onRefresh: () async{
                context.read<GetRobotsBloc>().add(GetRobot());
              },
              child: GridView.builder(
                gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                    crossAxisCount: 2,
                    crossAxisSpacing: 20,
                    mainAxisSpacing: 20,
                    childAspectRatio: 9 / 16),
                itemCount: state.robots.length,
                itemBuilder: (context, int i) {
                  log(state.robots[i].robotname ?? "error");
                  return Material(
                    elevation: 3,
                    color: Colors.blue[700],
                    shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(20)),
                    child: InkWell(
                      borderRadius: BorderRadius.circular(20),
                      onTap: () {
                        Navigator.push(
                            context,
                            MaterialPageRoute<void>(
                              builder: (BuildContext context) =>
                                  DetailsScreen(robot: state.robots[i]),
                            ));
                      },
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          Image.asset(
                            'lib/images/robot-test.png',
                          ),
                          Padding(
                            padding: const EdgeInsets.symmetric(
                                horizontal: 6, vertical: 6),
                            child: Row(
                              children: [
                                Container(
                                  decoration: BoxDecoration(
                                      color : (state.robots[i].robotstate ?? false)
                                      ? Colors.green.shade100
                                      : Colors.red.shade100,
                                      borderRadius: BorderRadius.circular(30)),
                                  child: Padding(
                                    padding: const EdgeInsets.symmetric(
                                        vertical: 4, horizontal: 4),
                                    child: Text(
                                      (state.robots[i].robotstate ?? false)
                                      ? "ONLINE"
                                      : "OFFLINE",
                                      style: TextStyle(
                                        color : (state.robots[i].robotstate ?? false)
                                        ? Colors.green
                                        : Colors.red,
                                        fontWeight: FontWeight.w800,
                                        fontSize: 8,
                                      ),
                                    ),
                                  ),
                                ),
                                const SizedBox(width: 8),
                                Container(
                                  decoration: BoxDecoration(
                                      color : (state.robots[i].robotidle ?? false)
                                      ? Colors.red.shade100
                                      : Colors.green.shade100,
                                      borderRadius: BorderRadius.circular(30)),
                                  child: Padding(
                                    padding: const EdgeInsets.symmetric(
                                        vertical: 4, horizontal: 4),
                                    child: Text(
                                      (state.robots[i].robotidle ?? false)
                                      ? "IDLE"
                                      : "NOT-IDLE",
                                      style: TextStyle(
                                        color : (state.robots[i].robotidle ?? false)
                                        ? Colors.red
                                        : Colors.green,
                                        fontWeight: FontWeight.w800,
                                        fontSize: 8,
                                      ),
                                    ),
                                  ),
                                )
                              ],
                            ),
                          ),
                          const SizedBox(height: 8),
                          Padding(
                            padding: const EdgeInsets.symmetric(horizontal: 12),
                            child: Text(
                              (state.robots[i].robotname ?? "error"),
                              style: const TextStyle(
                                fontSize: 14,
                                fontWeight: FontWeight.bold,
                              ),
                            ),
                          )
                        ],
                      ),
                    ),
                  );
                }));
            } else if(state is GetRobotLoading){
              return RefreshIndicator(
                child: const CircularProgressIndicator(),
                onRefresh: () async{
                  context.read<GetRobotsBloc>().add(GetRobot());
                },
              );
            } else {
              return RefreshIndicator(
                child: Center(child: Text("An error has occurred...")),
                onRefresh: () async{
                  context.read<GetRobotsBloc>().add(GetRobot());
                });
            }
          },
        ),
      ),
    );
  }
}
