
import 'package:flutter/cupertino.dart';
import 'package:robots_repository/robots_repository.dart' as macros;

class RobotEntity {
  String? iD;
  String? robotname;
  bool? robotidle;
  bool? robotstate;
  String? createdAt;
  String? updatedAt;
  String? robotId;
  String? userId;
  macros.MacrosEntity? robotInfo;
  List<macros.RouteEntity>? ruta;

  RobotEntity(
      {this.iD,
      this.robotname,
      this.robotidle,
      this.robotstate,
      this.createdAt,
      this.updatedAt,
      this.robotId,
      this.userId,
      this.robotInfo,
      this.ruta,});

  RobotEntity.fromJson(Map<String, dynamic> json) {
    iD = json['ID'];
    robotname = json['robotname'];
    robotidle = json['robotidle'];
    robotstate = json['robotstate'];
    createdAt = json['created_at'];
    updatedAt = json['updated_at'];
    robotId = json['robot_id'];
    userId = json['user_id'];
    robotInfo = json['robot_info'] != null
        ? macros.MacrosEntity.fromJson(json['robot_info'])
        : null;
    if (json['ruta'] != null) {
      ruta = <macros.RouteEntity>[];
      json['ruta'].forEach((v) {
        ruta!.add(macros.RouteEntity.fromJson(v));
      });
    }
  }

  get routes => null;

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = Map<String, dynamic>();
    data['ID'] = this.iD;
    data['robotname'] = this.robotname;
    data['robotidle'] = this.robotidle;
    data['robotstate'] = this.robotstate;
    data['created_at'] = this.createdAt;
    data['updated_at'] = this.updatedAt;
    data['robot_id'] = this.robotId;
    data['user_id'] = this.userId;
    if (this.robotInfo != null) {
      data['robot_info'] = this.robotInfo!.toJson();
    }
    if (this.ruta != null) {
      data['ruta'] = this.ruta!.map((v) => v.toJson()).toList();
    }
    return data;
  }
}