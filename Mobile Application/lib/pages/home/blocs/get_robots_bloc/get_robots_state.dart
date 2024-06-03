part of 'get_robots_bloc.dart';

@immutable
sealed class GetRobotsState {
  const GetRobotsState();

  @override
  List<Object> get props => [];
}

final class GetRobotsInitial extends GetRobotsState {}

final class GetRobotFailure extends GetRobotsState {}
final class GetRobotLoading extends GetRobotsState {}
final class GetRobotSuccess extends GetRobotsState {
  final List<RobotEntity> robots;

  const GetRobotSuccess(this.robots);

  @override
  List<Object> get props => [robots];

}

