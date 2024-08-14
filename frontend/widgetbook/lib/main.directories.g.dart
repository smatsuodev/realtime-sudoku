// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_import, prefer_relative_imports, directives_ordering

// GENERATED CODE - DO NOT MODIFY BY HAND

// **************************************************************************
// AppGenerator
// **************************************************************************

// ignore_for_file: no_leading_underscores_for_library_prefixes
import 'package:widgetbook/widgetbook.dart' as _i1;
import 'package:widgetbook_workspace/examples/cool_button.dart' as _i2;
import 'package:widgetbook_workspace/sudoku_game/sudoku_game.dart' as _i3;

final directories = <_i1.WidgetbookNode>[
  _i1.WidgetbookFolder(
    name: 'examples',
    children: [
      _i1.WidgetbookLeafComponent(
        name: 'NumberInputDialog(Answer mode)',
        useCase: _i1.WidgetbookUseCase(
          name: 'NumberInputDialog(Answer mode)',
          builder: _i3.buildNumberInputDialogAnswerUseCase,
        ),
      ),
      _i1.WidgetbookLeafComponent(
        name: 'NumberInputDialog(Note mode)',
        useCase: _i1.WidgetbookUseCase(
          name: 'NumberInputDialog(Note mode)',
          builder: _i3.buildNumberInputDialogNoteUseCase,
        ),
      ),
      _i1.WidgetbookLeafComponent(
        name: 'SudokuCell(Answer mode)',
        useCase: _i1.WidgetbookUseCase(
          name: 'SudokuCell(Answer mode)',
          builder: _i3.buildSudokuCellAnswerUseCase,
        ),
      ),
      _i1.WidgetbookLeafComponent(
        name: 'SudokuCell(Note mode)',
        useCase: _i1.WidgetbookUseCase(
          name: 'SudokuCell(Note mode)',
          builder: _i3.buildSudokuCellNoteUseCase,
        ),
      ),
      _i1.WidgetbookLeafComponent(
        name: 'SudokuGrid(Answer mode)',
        useCase: _i1.WidgetbookUseCase(
          name: 'SudokuGrid(Answer mode)',
          builder: _i3.buildSudokuGridAnswerUseCase,
        ),
      ),
      _i1.WidgetbookLeafComponent(
        name: 'SudokuGrid(Note mode)',
        useCase: _i1.WidgetbookUseCase(
          name: 'SudokuGrid(Note mode)',
          builder: _i3.buildSudokuGridNoteUseCase,
        ),
      ),
      _i1.WidgetbookLeafComponent(
        name: 'SudokuApp',
        useCase: _i1.WidgetbookUseCase(
          name: 'SudokuApp',
          builder: _i3.buildSudokuAppUseCase,
        ),
      ),
    ],
  )
];
