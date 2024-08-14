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

@widgetbook.UseCase(name: 'SudokuCell(Answer mode)', type: SudokuCell)
Widget buildSudokuCellAnswerUseCase(BuildContext context) {
  int value = 0;
  Set<int> notes = <int>{};
  bool isNoteMode = false;
  return SudokuCell(
    value: value,
    notes: notes,
    isNoteMode: isNoteMode,
    onChanged: (_) { return List.generate(9, (_) => <int>{});},
    onCleared: (_){},
    );
}


@widgetbook.UseCase(name: 'SudokuCell(Note mode)', type: SudokuCell)
Widget buildSudokuCellNoteUseCase(BuildContext context) {
  int value = 0;
  Set<int> notes = <int>{};
  bool isNoteMode = true;
  return SudokuCell(
    value: value,
    notes: notes,
    isNoteMode: isNoteMode,
    onChanged: (_) { return List.generate(9, (_) => <int>{});},
    onCleared: (_){},
    );
}
