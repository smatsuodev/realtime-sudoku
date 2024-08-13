import 'package:flutter/material.dart';
import 'package:frontend/presentation/page/rooms_list_page.dart';
import 'package:widgetbook_annotation/widgetbook_annotation.dart' as widgetbook;

@widgetbook.UseCase(name: 'Deafult', type: RoomsListPage)
Widget buildCoolButtonUseCase(BuildContext context) {
  return const RoomsListPage();
}
