import 'package:flutter/material.dart';
import 'package:widgetbook_annotation/widgetbook_annotation.dart' as widgetbook;
import 'package:frontend/presentation/component/progress_dialog.dart';

abstract class ProgressDialog {}

@widgetbook.UseCase(name: 'Default', type: ProgressDialog)
Widget buildProgressDialog(BuildContext context) {
  return Center(
    child: ElevatedButton(
      onPressed: () async {
        Future.delayed(const Duration(seconds: 2), () {
          hideProgressDialog(context);
        });
        showProgressDialog(context);
      },
      child: const Text("2秒後に消えるダイアログを表示"),
    ),
  );
}
