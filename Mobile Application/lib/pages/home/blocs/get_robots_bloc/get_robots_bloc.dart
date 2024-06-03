
import 'package:bloc/bloc.dart';
import 'package:meta/meta.dart';
import 'package:robots_repository/robots_repository.dart';

part 'get_robots_event.dart';
part 'get_robots_state.dart';

class GetRobotsBloc extends Bloc<GetRobotsEvent, GetRobotsState> {
  final RobotsRepo _robotsRepo;

  GetRobotsBloc(this._robotsRepo) : super(GetRobotsInitial()) {
    on<GetRobotsEvent>((event, emit) async {
      emit(GetRobotLoading());
      try{
        List<RobotEntity> robots = await _robotsRepo.getRobots();
        emit(GetRobotSuccess(robots));
      } catch (e) {
        emit(GetRobotFailure());
      }
    });
  }
}
