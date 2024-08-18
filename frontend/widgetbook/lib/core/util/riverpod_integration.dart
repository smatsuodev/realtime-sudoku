import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:widgetbook/widgetbook.dart';

class RiverpodIntegration extends WidgetbookIntegration {
  ProviderContainer get container {
    final overrides = _overrides;
    _overrides = []; // overridesを書き換えない時に意図しないoverridesによる挙動を防ぐため初期化しておく
    return ProviderContainer(overrides: overrides);
  }

  List<Override> _overrides = [];
  set overrides(List<Override> value) => _overrides = value;
}
