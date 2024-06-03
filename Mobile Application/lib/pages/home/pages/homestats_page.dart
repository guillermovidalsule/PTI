import 'package:flutter/material.dart';
import 'package:robots_repository/robots_repository.dart';
import 'package:talos/components/CircleProgress.dart';

class HomeStatsPage extends StatefulWidget{

  final RobotEntity robot;
  const HomeStatsPage({super.key, required this.robot});

  @override
  _HomeStatsState createState() => _HomeStatsState();

}

class _HomeStatsState extends State<HomeStatsPage> with SingleTickerProviderStateMixin{

  bool isLoading = false;
  late AnimationController progressController;
  late Animation<double> cpuFreqAnimation;
  late Animation<double> temperatureAnimation;
  

  @override
  void initState() {
    super.initState();

    double temp = widget.robot.robotInfo?.temperature ?? 0;
    double cpuFreq = widget.robot.robotInfo?.cpuFreq ?? 0;
    isLoading = true;
    _HomeStatsStateInit(temp, cpuFreq);
  }

  _HomeStatsStateInit(double temp, double cpuFreq){
    progressController = AnimationController(vsync: this, duration: Duration(milliseconds: 2000));

    temperatureAnimation = Tween<double>(begin: -50, end: temp).animate(progressController)..addListener(() {
      setState(() { });
    });

    cpuFreqAnimation = Tween<double>(begin: -20, end: cpuFreq).animate(progressController)..addListener(() {
      setState(() { });
    });

    progressController.forward();
  }

  @override
  Widget build(BuildContext context){
    return Scaffold(
      body: Center(
        child: isLoading ? Column(
          mainAxisAlignment: MainAxisAlignment.spaceEvenly,
          children: <Widget>[
            CustomPaint(
            foregroundPainter: CircleProgress(temperatureAnimation.value, true),
            child: SizedBox(
              width: 200,
              height: 200,
              child: Center(
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: <Widget>[
                    const Text('Temperature'),
                    Text(
                      '${temperatureAnimation.value.toInt()}',
                      style: const TextStyle(fontSize: 50, fontWeight: FontWeight.bold),
                    ),
                    const Text(
                      '°C',
                      style: TextStyle(
                        fontSize: 20,
                        fontWeight: FontWeight.bold
                      )
                    )
                  ]
                )
              )
            ),
          ),
          CustomPaint(
            foregroundPainter: CircleProgress(cpuFreqAnimation.value, false),
            child: Container(
              width: 200,
              height: 200,
              child: Center(
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: <Widget>[
                    const Text('CpuFreq'),
                    Text(
                      '${cpuFreqAnimation.value.toInt()}',
                      style: const TextStyle(
                        fontSize: 50, fontWeight: FontWeight.bold
                      ),
                    )
                  ]
                )
              )
            )
          )
          
          ],
        ) : const Text (
          'Loading...',
          style: TextStyle(
            fontSize: 20,
            fontWeight: FontWeight.bold
          ),
        )
      ),
    );
  }


}