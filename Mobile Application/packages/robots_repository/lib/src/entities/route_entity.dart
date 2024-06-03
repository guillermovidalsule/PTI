class RouteEntity {
  double? latitud;
  double? longitud;

  RouteEntity({this.latitud, this.longitud});

  RouteEntity.fromJson(Map<String, dynamic> json) {
    latitud = json['latitud'];
    longitud = json['longitud'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['latitud'] = latitud;
    data['longitud'] = longitud;
    return data;
  }
}