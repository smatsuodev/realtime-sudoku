import 'package:flutter/material.dart';
import 'package:widgetbook_annotation/widgetbook_annotation.dart' as widgetbook;

class CoolButton extends StatelessWidget {
  const CoolButton({super.key});

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      onPressed: () {},
      child: const Text('Cool Button'),
    );
  }
}

@widgetbook.UseCase(name: 'Deafult', type: CoolButton)
Widget buildCoolButtonUseCase(BuildContext context) {
  return const CoolButton();
}
