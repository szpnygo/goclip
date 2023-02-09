import 'package:goclip/app/ui/pages/main/main_binding.dart';
import 'package:goclip/app/ui/pages/main/main_page.dart';
import 'package:get/get.dart';

part './app_routes.dart';

abstract class AppPages {
  static final pages = [
    GetPage(
      name: Routes.INITIAL,
      page: () => const MainPage(),
      binding: MainBinding(),
    ),
  ];
}
