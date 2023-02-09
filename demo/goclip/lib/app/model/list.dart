import 'package:json_annotation/json_annotation.dart';

part 'list.g.dart';

@JsonSerializable()
class ImageList {
  String? token;
  List<String>? images;

  ImageList({this.token, this.images});

  factory ImageList.fromJson(Map<String, dynamic> json) =>
      _$ImageListFromJson(json);

  Map<String, dynamic> toJson() => _$ImageListToJson(this);
}
