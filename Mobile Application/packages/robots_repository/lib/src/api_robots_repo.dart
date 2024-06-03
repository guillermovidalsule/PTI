import 'dart:convert';
import 'dart:developer';

import 'package:http/http.dart' as http;
import 'package:robots_repository/robots_repository.dart';
import 'package:shared_preferences/shared_preferences.dart';

class ApiRobotsRepo implements RobotsRepo {


  Future<List<RobotEntity>> getUserRobots() async {
    final prefs = await SharedPreferences.getInstance();
    List<RobotEntity> robots = List.empty();
    var headers  = {'Content-Type': 'application/json', 'token': prefs.get('token').toString()};
    try {
      var url = Uri.parse("http://nattech.fib.upc.edu:40342/robot/listar");
      http.Response response = 
        await http.get(url, headers: headers);

      if(response.contentLength == 0 || response.body == "null"){
        return List.empty();
      } else {
        final List parsedList = jsonDecode(response.body);
        log(parsedList.toString());
        List<RobotEntity> map = (parsedList).map((i) => RobotEntity.fromJson(i)).toList();
        return map;
      }
    } catch (e){
      log(e.toString());
      rethrow;
    }
  }

  @override
  Future<List<RobotEntity>> getRobots() async {
    try{
      return getUserRobots();
    } catch (e) {
      log(e.toString());
      rethrow;
    }
  }
}