import 'package:flutter_modular/flutter_modular.dart';
import 'package:frontend/module/dashboard/presentation/pages/approval_success.dart';
import 'package:frontend/module/dashboard/presentation/pages/approval_token-2.dart';
import 'package:frontend/module/dashboard/presentation/pages/approval_token.dart';
import 'package:frontend/module/dashboard/presentation/pages/dashboard_selected.dart';
import 'package:frontend/module/dashboard/presentation/pages/home_page.dart';

class AppModule extends Module {

  @override
  List<Bind> get binds => [

  ];

  @override
  List<ChildRoute> get routes => [
    ChildRoute('/', child: (context, args) => const HomePage()),
    ChildRoute('/dashboard-selected', child: (context, args) => const DashboardSelected()),
    ChildRoute('/token-1', child:(context, args) => TokenApprovalScreen() ),
    ChildRoute('/token-2', child:(context, args) => TokenApprovalTwoScreen() ),
    ChildRoute('/token-3', child: (context, args) => const ApprovalSuccess())
  ];
}