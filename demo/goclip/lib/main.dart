import 'package:goclip/app/api/images.dart';
import 'package:goclip/app/routes/app_pages.dart';
import 'package:goclip/app/services/image_service.dart';
import 'package:goclip/app/ui/pages/main/main_binding.dart';
import 'package:goclip/app/ui/pages/main/main_page.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:get/get.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await initServices();
  runApp(GetMaterialApp(
    title: "goclip-demo",
    debugShowCheckedModeBanner: false,
    initialRoute: Routes.INITIAL,
    home: const MainPage(),
    initialBinding: MainBinding(),
    defaultTransition: Transition.fadeIn,
    getPages: AppPages.pages,
    builder: EasyLoading.init(),
  ));
}

Future<void> initServices() async {
  Get.lazyPut(() => ImagesProvider());
  Get.lazyPut(() => ImageService());
}
