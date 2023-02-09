import 'package:goclip/app/api/base/http.dart';
import 'package:goclip/app/model/result.dart';

class BaseProvider {
  HttpHelper helper = HttpHelper();

  Future<Result<T>> getx<T>(
    String url, {
    Map<String, dynamic>? queryParameters,
  }) {
    return helper.get(url, queryParameters: queryParameters);
  }

  Future<Result<T>> postx<T>(String url, String data) async {
    return helper.post(url, data);
  }

  Future<Result<T>> uploadx<T>(
      String url, List<int> data, String fileName) async {
    return helper.upload(url, data, fileName);
  }
}
