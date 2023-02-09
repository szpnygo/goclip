import 'package:goclip/app/services/image_service.dart';
import 'package:file_picker/_internal/file_picker_web.dart';
import 'package:file_picker/file_picker.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:get/get.dart';

class MainController extends GetxController {
  FocusNode focusNode = FocusNode();
  ScrollController scrollController = ScrollController();
  ImageService imageService = Get.find();

  final _imageList = <String>[].obs;
  List<String> get imageList => _imageList;
  set imageList(List<String> value) => _imageList.value = value;

  MainController() {
    scrollController.addListener(() {
      if (scrollController.position.pixels ==
          scrollController.position.maxScrollExtent) {}
    });

    loadImages("");
  }

  loadImages(String token) async {
    if (token.isEmpty) {
      imageList.clear();
    }
    var result = await imageService.list(token);
    if (result.isSuccess()) {
      for (var element in result.data?.images ?? []) {
        imageList.add(element);
      }
    }
  }

  uploadImage() async {
    EasyLoading.show(status: "uploading...");
    FilePickerResult? result = await FilePickerWeb.platform.pickFiles(
      allowedExtensions: ['jpg', 'png', 'jpeg'],
      type: FileType.custom,
    );
    if (result != null) {
      if (result.files.single.bytes != null) {
        var uploadResult = await imageService.upload(
          result.files.single.bytes!,
          result.files.single.name.replaceAll(" ", ""),
        );
        if (uploadResult.isSuccess()) {
          loadImages("");
        } else {
          EasyLoading.showError("upload image failed");
        }
      }
    }
    EasyLoading.dismiss();
  }

  search(String data) async {
    if (data.isEmpty) {
      loadImages("");
      return;
    }

    var result = await imageService.search(data);
    imageList.clear();
    if (result.isSuccess()) {
      for (var element in result.data ?? []) {
        imageList.add(element);
      }
    }
  }
}
