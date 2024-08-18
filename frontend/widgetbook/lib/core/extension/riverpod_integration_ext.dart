import 'package:widgetbook/widgetbook.dart';
import 'package:widgetbook_workspace/core/util/riverpod_integration.dart';

extension RiverpodIntegrationExtension on WidgetbookState {
  RiverpodIntegration get riverpodIntegration =>
      integrations?.firstWhere((element) => element is RiverpodIntegration)
          as RiverpodIntegration;
}
