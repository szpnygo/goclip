import 'dart:ui';

import 'package:goclip/app/ui/pages/main/main_controller.dart';
import 'package:flutter/material.dart';
import 'package:flutter_staggered_grid_view/flutter_staggered_grid_view.dart';
import 'package:get/get.dart';

class MainPage extends GetView<MainController> {
  const MainPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          controller.uploadImage();
        },
        child: const Icon(Icons.upload_file),
      ),
      backgroundColor: const Color(0xff282828),
      body: SafeArea(
        child: body(context),
      ),
    );
  }

  body(context) {
    return Column(
      children: [
        const Padding(
          padding: EdgeInsets.only(top: 24),
          child: Center(
            child: Text(
              "此项目仅供测试使用，请勿上传任务私人图片或违反法律法规之图片。同时请勿滥用该服务，本demo随时可能关闭。\nThis project is for test use only, please do not upload private pictures of tasks or pictures that violate laws and regulations. At the same time, please do not abuse the service, this demo may be closed at any time.\n",
              style: TextStyle(color: Colors.white),
              textAlign: TextAlign.center,
            ),
          ),
        ),
        searchBar(),
        Expanded(
          child: Padding(
            padding: const EdgeInsets.only(
              left: 32,
              right: 32,
              top: 16,
              bottom: 32,
            ),
            child: imageGallery(context),
          ),
        ),
      ],
    );
  }

  searchBar() {
    return Padding(
      padding: const EdgeInsets.only(left: 32, right: 32, top: 12),
      child: TextField(
        style: const TextStyle(
          color: Colors.white,
        ),
        focusNode: controller.focusNode,
        onSubmitted: (value) {
          controller.focusNode.requestFocus();
          controller.search(value);
        },
        decoration: InputDecoration(
          hintText: 'Search -- Currently only supports English',
          hintStyle: const TextStyle(
            color: Colors.grey,
          ),
          prefixIcon: const Icon(
            Icons.search,
            color: Colors.white,
          ),
          border: OutlineInputBorder(
            borderRadius: BorderRadius.circular(10),
            borderSide: BorderSide.none,
          ),
          filled: true,
          fillColor: Colors.grey[800],
        ),
      ),
    );
  }

  imageGallery(context) {
    return ScrollConfiguration(
      behavior: ScrollConfiguration.of(context).copyWith(
        dragDevices: {
          PointerDeviceKind.touch,
          PointerDeviceKind.mouse,
        },
        scrollbars: false,
      ),
      child: SingleChildScrollView(
        controller: controller.scrollController,
        child: Obx(
          () => StaggeredGrid.count(
            crossAxisCount: 5,
            mainAxisSpacing: 10,
            crossAxisSpacing: 10,
            children: controller.imageList
                .map(
                  (e) => Image.network(
                    "https://goclip-1300434835.picsh.myqcloud.com/$e?imageMogr2/thumbnail/x512/format/webp",
                    fit: BoxFit.cover,
                    height: 360,
                  ),
                )
                .toList(),
          ),
        ),
      ),
    );
  }
}
