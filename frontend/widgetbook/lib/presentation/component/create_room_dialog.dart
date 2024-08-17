import 'package:flutter/material.dart';
import 'package:frontend/presentation/component/create_room_dialog.dart';
import 'package:widgetbook_annotation/widgetbook_annotation.dart' as widgetbook;

@widgetbook.UseCase(name: 'Default', type: CreateRoomDialog)
Widget buildCreateRoomDialog(BuildContext context) {
  return Center(
    child: ElevatedButton(
      onPressed: () {
        showDialog(
          context: context,
          builder: (context) => CreateRoomDialog(
            onSubmit: (name) {
              print('Room name: $name');
            },
          ),
        );
      },
      child: const Text("Open Dialog"),
    ),
  );
}
