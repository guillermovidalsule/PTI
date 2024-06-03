import 'package:robots_repository/robots_repository.dart';

abstract class RobotsRepo {
  Future<List<RobotEntity>> getRobots();
}