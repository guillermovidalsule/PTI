import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:robots_repository/robots_repository.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:talos/pages/home/blocs/get_robots_bloc/get_robots_bloc.dart';
import 'package:talos/pages/home/home_page.dart';
import 'package:talos/pages/auth/login_or_register_page.dart';

class AuthPage extends StatelessWidget {
  const AuthPage({super.key});

  @override
  Widget build(BuildContext context){
    final stream = Stream.periodic(const Duration(seconds: 0)).asyncMap((_) => SharedPreferences.getInstance());
    return StreamBuilder<SharedPreferences>(
      key: ValueKey<Stream<SharedPreferences>>(stream),
      stream: stream,
      builder: (context, snapshot){
        print('Building Stream');
        if(snapshot.hasData){
          final prefs = snapshot.data!;
          final token  = prefs.getString('token');
          if(token != null){
            //return HomePage();
            return BlocProvider(
              create: (context) => GetRobotsBloc(
                ApiRobotsRepo()
              )..add(GetRobot()),
              child: const HomePage(),
            );
          } else {
            return const LoginOrRegisterPage();
          }
        } else {
          return const CircularProgressIndicator();
        }
      },

    );
  }
}