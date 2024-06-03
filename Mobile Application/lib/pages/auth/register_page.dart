import 'dart:convert';
import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:http/http.dart' as http;
import 'package:talos/components/my_button.dart';
import 'package:talos/components/my_textfield.dart';
import 'package:talos/utils/api_endpoints.dart';
import 'package:shared_preferences/shared_preferences.dart';

class RegisterPage extends StatefulWidget {
  final Function()? onTap;
  const RegisterPage({super.key, required this.onTap});

  @override
  State<RegisterPage> createState() => _RegisterPageState();
}

class _RegisterPageState extends State<RegisterPage> {
  final usernameController = TextEditingController();

  final passwordController = TextEditingController();

  final emailController = TextEditingController();

  final confirmPasswordController = TextEditingController();

  final Future<SharedPreferences> _prefs = SharedPreferences.getInstance();

  void error(){
    Get.back();
    showDialog(
        context: Get.context!,
        builder: (context) {
          return const SimpleDialog(
            title: Text('Error'),
            contentPadding: EdgeInsets.all(20),
            children: [Text("ERROR ELSE LOGIN")],
          );
        });
  }

  void showErrorMessage(String message) {
    showDialog(
      context: context,
      builder: (context) {
        return AlertDialog(
          title: Text(message),
        );
      }
    );
  }

  void register() async {
    //Circular progression bar
    showDialog(
        context: context,
        builder: (context) {
          return const Center(
            child: CircularProgressIndicator(),
          );
        });

    var headers = {'Content-Type': 'application/json'};
    try {
      if (passwordController.text == confirmPasswordController.text) {
        var url = Uri.parse(
            ApiEndPoints.baseUrl + ApiEndPoints.authEndPoints.register);
        Map body = {
          'email': emailController.text,
          'username': usernameController.text,
          'password': passwordController.text
        };
        http.Response response =
            await http.post(url, body: jsonEncode(body), headers: headers);
        Navigator.pop(context);

        log(response.body);
        final json = jsonDecode(response.body);
        if (response.statusCode != 200) {
            if(json['error'] == "Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag"){
              showErrorMessage("Wrong email provided");
            }
            if(json['error'] == "Key: 'User.Password' Error:Field validation for 'Password' failed on the 'min' tag"){
              showErrorMessage("Password must be minimun 8 characters");
            }
            if(json['error'] == "User item was not created"){
              showErrorMessage("Username or email already exists");
            }
        } else {
          var token = json['token'];
          final SharedPreferences prefs = await _prefs;
          await prefs.setString('token', token);

          usernameController.clear();
          emailController.clear();
          passwordController.clear();
          confirmPasswordController.clear();
          //Get.off(() => const HomePage());
        }
      } else {
        showErrorMessage("Passwords don't match");
      }
    } catch (e) {
      Get.back();
      showDialog(
          context: Get.context!,
          builder: (context) {
            return SimpleDialog(
              title: const Text('Error'),
              contentPadding: const EdgeInsets.all(20),
              children: [Text(e.toString())],
            );
          });
    }
  }

  @override
  Widget build(BuildContext context){
    return Scaffold(
      backgroundColor: Colors.lightBlue[50],
      body: SafeArea(
        child: Center(
          child: SingleChildScrollView(
            child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Image.asset('lib/images/logo.png', width: 200, height: 150),
              
              const SizedBox(height: 10),
            
              MyTextField(
                controller: emailController,
                hintText: 'Email',
                obscureText: false,
              ),
            
              const SizedBox(height: 20),
            
              MyTextField(
                controller: usernameController,
                hintText: 'Username',
                obscureText: false,
              ),

              const SizedBox(height: 20),
            
              MyTextField(
                controller: passwordController,
                hintText: 'Password',
                obscureText: true,
              ),

              const SizedBox(height: 25),
            
              MyTextField(
                controller: confirmPasswordController,
                hintText: 'Confirm Password',
                obscureText: true,
              ),
            
              const SizedBox(height: 25),
            
              MyButton(
                onTap: register,
                text: "Register",
              ),
            
              const SizedBox(height: 50),
            
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text('Already have an account?', style: TextStyle(color: Colors.grey[500]),),
                  const SizedBox(width: 4),
                  GestureDetector(
                    onTap: widget.onTap,
                    child: const Text('Login Now', style: TextStyle(color: Colors.blue, fontWeight: FontWeight.bold),
                    ),
                  ),
                ],
              )
            ],),
          ),
        ),
      )
    );
  }
}