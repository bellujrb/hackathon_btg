import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:frontend/core/extensions/build_context_utils.dart';

class ApprovalSuccess extends StatelessWidget {
  const ApprovalSuccess({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: InkWell(
          onTap: () => Modular.to.navigate("/"),
          child: SizedBox(
            height: context.mediaHeight * 0.6,
            child: Image.asset('assets/btg3.png'),
          ),
        ),
      ),
    );
  }
}