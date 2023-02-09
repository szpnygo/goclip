import 'dart:convert';

import 'package:goclip/app/api/base/base.dart';
import 'package:goclip/app/model/list.dart';
import 'package:goclip/app/model/result.dart';

class ImagesProvider extends BaseProvider {
  // Get the list of images
  Future<Result<ImageList>> list(String token) {
    var request = {
      "token": token,
      "size": 50,
    };
    return postx<ImageList>("/images", jsonEncode(request));
  }

  Future<Result<List<String>>> search(String data) {
    return getx<List<String>>("/search?search=$data");
  }

  Future<Result<void>> upload(List<int> data, String fileName) {
    return uploadx("/image", data, fileName);
  }
}
