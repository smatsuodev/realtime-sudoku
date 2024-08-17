import 'package:flutter/material.dart';
import 'package:frontend/presentation/component/create_room_button.dart';
import 'package:widgetbook_annotation/widgetbook_annotation.dart' as widgetbook;

@widgetbook.UseCase(name: 'Deafult', type: CreateRoomButton)
Widget buildCreateRoomButton(BuildContext context) {
  return const CreateRoomButton();
}
