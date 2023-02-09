import 'package:goclip/app/api/images.dart';
import 'package:goclip/app/model/list.dart';
import 'package:goclip/app/model/result.dart';
import 'package:get/get.dart';

class ImageService extends GetxService {
  late ImagesProvider provider = Get.find();

  Future<ImageService> init() async {
    return this;
  }

  Future<Result<ImageList>> list(String token) {
    return provider.list(token);
  }

  Future<Result<void>> upload(List<int> data, String fileName) {
    return provider.upload(data, fileName);
  }

  Future<Result<List<String>>> search(String data) {
    return provider.search(data);
  }
}
