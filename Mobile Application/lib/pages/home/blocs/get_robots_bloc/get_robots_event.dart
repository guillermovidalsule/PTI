part of 'get_robots_bloc.dart';

@immutable
sealed class GetRobotsEvent {
  const GetRobotsEvent();

  @override
  List<Object> get props => [];
}

class GetRobot extends GetRobotsEvent{}
