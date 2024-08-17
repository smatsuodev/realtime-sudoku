import 'package:flutter/material.dart';
import 'package:frontend/domain/entity/room.dart';
import 'package:frontend/presentation/page/rooms_list_page.dart';
import 'package:widgetbook_annotation/widgetbook_annotation.dart' as widgetbook;

@widgetbook.UseCase(name: 'Deafult', type: RoomsListPage)
Widget buildCoolButtonUseCase(BuildContext context) {
  final rooms = List.generate(
    16,
    (index) => Room(
      id: index,
      name: 'Room $index',
      players: [],
    ),
  ).toList();
  return RoomsListPage(rooms: rooms);
}
