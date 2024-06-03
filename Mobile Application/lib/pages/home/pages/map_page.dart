
import 'dart:math';

import 'package:flutter/material.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:robots_repository/robots_repository.dart';

class MapPage extends StatefulWidget {
  final RobotEntity robot;
  const MapPage({super.key, required this.robot});

  @override
  State<MapPage> createState() => _MapPageState();
}

class _MapPageState extends State<MapPage> {

  final Set<Marker> markers = new Set();

  static const LatLng position = const LatLng(41.389486, 2.113405);
  bool isNull = false;

  @override
  void initState(){
    super.initState();
    int? length = widget.robot.ruta?.length;
    if(widget.robot.ruta == null){
      isNull = true;
    } else {
      for(int i = 0; i < length!; ++i){
        double latitud = widget.robot.ruta?[i].latitud ?? 0.0;
        double longitud = widget.robot.ruta?[i].longitud ?? 0.0;
        log(longitud);
        markers.add(Marker(
          markerId: MarkerId('test${i}'),
          position: LatLng(latitud, longitud),
          icon: BitmapDescriptor.defaultMarker
        ));
      }
    }
  }

  @override
  Widget build(BuildContext context){
    return Scaffold(
      body: (isNull == false)
       ? GoogleMap(
        initialCameraPosition: const CameraPosition(
          target: position,
          zoom: 15
        ),
        markers: markers
      )
      : const Center(
        child: Text("Aquest robot no té una ruta seleccionada"),
      )
    );
  }
}