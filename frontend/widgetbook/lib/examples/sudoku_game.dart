import 'package:flutter/material.dart';
import 'package:frontend/page/sudoku_game.dart';
import 'package:widgetbook_annotation/widgetbook_annotation.dart' as widgetbook;

@widgetbook.UseCase(name: 'NumberInputDialog(Answer mode)', type: NumberInputDialog)
Widget buildNumberInputDialogAnswerUseCase(BuildContext context) {
  return NumberInputDialog(isNoteMode: false);
}

@widgetbook.UseCase(name: 'NumberInputDialog(Note mode)', type: NumberInputDialog)
Widget buildNumberInputDialogNoteUseCase(BuildContext context) {
  return NumberInputDialog(isNoteMode: true);
}
