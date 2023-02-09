import 'package:goclip/app/config/constant.dart';
import 'package:goclip/app/model/result.dart';
import 'package:dio/dio.dart';

class HttpHelper {
  late Dio dioClient;

  HttpHelper() {
    dioClient = Dio();
    dioClient.options.baseUrl = Constants.baseAPI;
  }

  Future<Result<T>> get<T>(
    String url, {
    Map<String, dynamic>? queryParameters,
  }) async {
    try {
      var result = await dioClient.get(url, queryParameters: queryParameters);
      if (result.statusCode == 200) {
        return Future.value(Result<T>.fromJson(result.data));
      }
      return Future.value(
        Result(code: 1, message: "网络请求失败，请重试", data: null),
      );
    } on DioError catch (e) {
      if (e.response != null) {
        return Future.value(
          Result(code: 1, message: e.response!.statusMessage!, data: null),
        );
      } else {
        return Future.value(
          Result(code: 1, message: e.message, data: null),
        );
      }
    }
  }

  Future<Result<T>> post<T>(String url, String data) async {
    try {
      var result = await dioClient.post(url, data: data);
      if (result.statusCode == 200) {
        return Future.value(Result<T>.fromJson(result.data));
      }
      return Future.value(
        Result(code: 1, message: "网络请求失败，请重试", data: null),
      );
    } on DioError catch (e) {
      if (e.response != null) {
        return Future.value(
          Result(code: 1, message: e.response!.statusMessage!, data: null),
        );
      } else {
        return Future.value(
          Result(code: 1, message: e.message, data: null),
        );
      }
    }
  }

  Future<Result<T>> upload<T>(
      String url, List<int> value, String fileName) async {
    var data = FormData.fromMap({
      "file": MultipartFile.fromBytes(value, filename: fileName),
    });
    try {
      var result = await dioClient.post(url, data: data);
      if (result.statusCode == 200) {
        return Future.value(Result<T>.fromJson(result.data));
      }
      return Future.value(
        Result(code: 1, message: "网络请求失败，请重试", data: null),
      );
    } on DioError catch (e) {
      if (e.response != null) {
        return Future.value(
          Result(code: 1, message: e.response!.statusMessage!, data: null),
        );
      } else {
        return Future.value(
          Result(code: 1, message: e.message, data: null),
        );
      }
    }
  }
}
