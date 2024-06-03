class ApiEndPoints {
  static const String baseUrl = 'http://nattech.fib.upc.edu:40342';
  static _AuthEndPoints authEndPoints = _AuthEndPoints();
}

class _AuthEndPoints {
  final String register = '/users/signup';
  final String login = '/users/login';
  final String addRobot = '/robot/alta';
  final String removeRobot = '/robot/baja/';
  final String listRobots = '/robot/listar';
}