class MacrosEntity {
  double? cpuFreq;
  double? temperature;
  double? velocity;

  MacrosEntity({this.cpuFreq, this.temperature, this.velocity});

  MacrosEntity.fromJson(Map<String, dynamic> json) {
    cpuFreq = json['cpu_freq'].toDouble();
    temperature = json['temperature'].toDouble();
    velocity = json['velocity'].toDouble();
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['cpu_freq'] = cpuFreq;
    data['temperature'] = temperature;
    data['velocity'] = velocity;
    return data;
  }
}