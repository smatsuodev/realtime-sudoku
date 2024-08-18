import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:frontend/presentation/navigation/dialog_navigation.dart';

class CreateRoomDialog extends ConsumerStatefulWidget {
  final Function(String name) onSubmit;
  const CreateRoomDialog({super.key, required this.onSubmit});

  @override
  ConsumerState<CreateRoomDialog> createState() => _CreateRoomDialogState();
}

class _CreateRoomDialogState extends ConsumerState<CreateRoomDialog> {
  String name = "";

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: const Text("ルームを新規作成"),
      insetPadding: const EdgeInsets.all(20),
      content: TextField(
        decoration: const InputDecoration(
          labelText: 'ルーム名',
        ),
        onChanged: (value) {
          setState(() {
            name = value;
          });
        },
        autofocus: true,
      ),
      actions: [
        TextButton(
          onPressed: () {
            DialogNavigation.pop(context);
          },
          child: const Text('キャンセル'),
        ),
        TextButton(
          onPressed: () {
            widget.onSubmit(name);
            DialogNavigation.pop(context);
          },
          child: const Text('作成'),
        ),
      ],
    );
  }
}
