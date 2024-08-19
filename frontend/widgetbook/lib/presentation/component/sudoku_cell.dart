import 'package:flutter/material.dart';
import 'package:frontend/presentation/component/sudoku_cell.dart';
import 'package:widgetbook_annotation/widgetbook_annotation.dart' as widgetbook;

class _SudokuCellWrapper extends StatefulWidget {
  @override
  State<_SudokuCellWrapper> createState() => _SudokuCellWrapperState();
}

class _SudokuCellWrapperState extends State<_SudokuCellWrapper> {
  List<List<int>> cellVelues =
      List.generate(2, (_) => List.generate(2, (_) => 0));

  @override
  Widget build(BuildContext context) {
    final children = cellVelues.indexed.map((t) {
        final (y, row) = t;
    return Row(
        children: 
            row.indexed.map((t) {
                final (x, value) = t;
          return SudokuCell(
            value: value,
            onChanged: (newValue) {
              setState(() {
                cellVelues[y][x] = newValue;
              });
            },
          );
        }).toList()
    );
    }).toList();
    return Column(
        children: children);
  }
}

@widgetbook.UseCase(name: 'Default', type: SudokuCell)
Widget buildSudokuCell(BuildContext context) {
  return Center(
      child: _SudokuCellWrapper());
}
