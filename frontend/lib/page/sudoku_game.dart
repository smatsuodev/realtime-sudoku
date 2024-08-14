import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
class SudokuApp extends StatefulWidget {

class SudokuCell extends StatelessWidget {
  final int value;
  final Set<int> notes;
  final bool isNoteMode;
  final Function(List<int>) onChanged;
  final Function onCleared;

  SudokuCell({required this.value, required this.notes, required this.isNoteMode, required this.onChanged, required this.onCleared});

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(1.0),
      child: GestureDetector(
        onTap: () async {
          String? input = await showDialog<String>(
            context: context,
            builder: (context) => NumberInputDialog(isNoteMode: isNoteMode),
          );
          if (input != null) {
            input = input.replaceAll(' ', ''); // Remove spaces
            if (input.isEmpty) {
              onCleared(); // Clear the cell and restore notes
            } else {
              List<int> newValues = input.split('').map(int.parse).where((n) => n >= 1 && n <= 9).toList();
              if (newValues.isNotEmpty) {
                onChanged(newValues);
              }
            }
          }
        },
        child: Container(
          decoration: BoxDecoration(
            border: Border.all(color: Colors.black),
          ),
          child: Center(
            child: value == 0 && notes.isNotEmpty
                ? Text(
                    notes.map((e) => e.toString()).join(', '),
                    textAlign: TextAlign.center,
                    style: TextStyle(fontSize: 10, color: Colors.grey),
                  )
                : Text(
                    value == 0 ? '' : value.toString(),
                    textAlign: TextAlign.center,
                    style: TextStyle(fontSize: 24),
                  ),
          ),
        ),
      ),
    );
  }
}

class NumberInputDialog extends StatelessWidget {
  final bool isNoteMode;

  NumberInputDialog({required this.isNoteMode});

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: Text(isNoteMode ? 'Enter Numbers for Notes' : 'Enter Number'),
      content: TextField(
        keyboardType: TextInputType.text,
        decoration: InputDecoration(hintText: isNoteMode ? 'Enter numbers (1-9) separated by spaces or without spaces' : 'Enter a number (1-9) or leave empty to clear'),
        onSubmitted: (value) {
          if (Navigator.of(context).maybePop(value) == true) {
            Navigator.of(context).pop(value);
          }
        },
      ),
    );
  }
}
