import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:widgetbook/widgetbook.dart';
import 'package:widgetbook_annotation/widgetbook_annotation.dart' as widgetbook;
import 'package:widgetbook_workspace/core/extension/riverpod_integration_ext.dart';
import 'package:widgetbook_workspace/core/util/riverpod_integration.dart';

// This file does not exist yet,
// it will be generated in the next step
import 'main.directories.g.dart';

void main() {
  runApp(const WidgetbookApp());
}

@widgetbook.App()
class WidgetbookApp extends StatelessWidget {
  const WidgetbookApp({super.key});

  @override
  Widget build(BuildContext context) {
    return Widgetbook.material(
      directories: directories,
      appBuilder: (context, child) => UncontrolledProviderScope(
        container: WidgetbookState.of(context).riverpodIntegration.container,
        child: MaterialApp(home: child),
      ),
      integrations: [
        RiverpodIntegration(),
      ],
    );
  }
}
