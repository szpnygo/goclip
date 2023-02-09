import 'package:goclip/app/model/list.dart';
import 'package:json_annotation/json_annotation.dart';

class Result<T> {
  final int code;
  final String message;

  @JsonKey(fromJson: _dataFromJson)
  final T? data;

  Result({required this.code, required this.message, this.data});

  bool isSuccess() {
    return code == 0;
  }

  @override
  String toString() {
    return 'Result{code: $code, message: $message, data: $data}';
  }

  factory Result.fromJson(Map<String, dynamic> json) {
    if (T.toString() == "void") {
      return Result<T>(
        code: json['code'] as int,
        message: json['message'] as String,
        data: null,
      );
    }
    if (json['data'] == null) {
      return Result<T>(
        code: json['code'] as int,
        message: json['message'] as String,
        data: null,
      );
    }
    return Result<T>(
      code: json['code'] as int,
      message: json['message'] as String,
      data: Result._dataFromJson(json['data'] as Object),
    );
  }

  static T _dataFromJson<T>(Object json) {
    if (json is List) {
      if (T is List<String> ||
          T.toString() == "List<String>" ||
          T.toString() == "List<String>?" ||
          T.toString() == "${<String>[].runtimeType}?") {
        return json.map((e) => e as String).toList() as T;
      }
    }
    if (json is Map<String, dynamic>) {
      if (T is ImageList ||
          T.toString() == "ImageList" ||
          T.toString() == "ImageList?" ||
          T.toString() == "${ImageList().runtimeType}?") {
        return ImageList.fromJson(json) as T;
      }
    }

    throw ArgumentError.value(
      json,
      'json',
      'Cannot convert the provided data.',
    );
  }
}
