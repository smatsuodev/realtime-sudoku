import 'package:flutter/material.dart';
import 'package:frontend/domain/repository/room_repository.dart';
import 'package:frontend/presentation/page/rooms_list_page.dart';
import 'package:widgetbook/widgetbook.dart';
import 'package:widgetbook_annotation/widgetbook_annotation.dart' as widgetbook;
import 'package:widgetbook_workspace/core/extension/riverpod_integration_ext.dart';
import 'package:widgetbook_workspace/infrastructure/repository/room_repository_impl.dart';

@widgetbook.UseCase(name: 'Default', type: RoomsListPage)
Widget buildRoomsListPage(BuildContext context) {
  final shouldDelay = context.knobs.boolean(label: 'APIの遅延');
  WidgetbookState.of(context).riverpodIntegration.overrides = [
    roomRepositoryProvider.overrideWithValue(MockRoomRepositoryImpl(
      shouldDelay: shouldDelay,
    )),
  ];
  return RoomsListPage(key: UniqueKey());
}
