import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

class SudokuApp extends StatefulWidget {
  @override
  _SudokuAppState createState() => _SudokuAppState();
}


class _SudokuAppState extends State<SudokuApp> {
  bool isNoteMode = false;

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: Text('Sudoku'),
          actions: [
            IconButton(
              icon: Icon(isNoteMode ? Icons.note : Icons.edit),
              onPressed: toggleMode,
            ),
          ],
        ),
        body: Shortcuts(
          shortcuts: {
            LogicalKeySet(LogicalKeyboardKey.space): ToggleModeIntent(),
          },
          child: Actions(
            actions: {
              ToggleModeIntent: CallbackAction<ToggleModeIntent>(
                onInvoke: (intent) => toggleMode(),
              ),
            },
            child: Focus(
              autofocus: true,
              child: SudokuGrid(isNoteMode: isNoteMode),
            ),
          ),
        ),
      ),
    );
  }

  void toggleMode() {
    setState(() {
      isNoteMode = !isNoteMode;
    });
  }
}

class ToggleModeIntent extends Intent {}

class SudokuGrid extends StatefulWidget {
  final bool isNoteMode;

  SudokuGrid({required this.isNoteMode});

  @override
  _SudokuGridState createState() => _SudokuGridState();
}

class _SudokuGridState extends State<SudokuGrid> {
  final List<List<int>> sudokuBoard = List.generate(9, (_) => List.generate(9, (_) => 0));
  final List<List<Set<int>>> notes = List.generate(9, (_) => List.generate(9, (_) => <int>{}));

  @override
  Widget build(BuildContext context) {
    double screenWidth = MediaQuery.of(context).size.width;
    double screenHeight = MediaQuery.of(context).size.height;
    double availableHeight = screenHeight - kToolbarHeight - MediaQuery.of(context).padding.top - 16.0;
    double gridSize = screenWidth < availableHeight ? screenWidth : availableHeight;

    return Center(
      child: Container(
        width: gridSize,
        height: gridSize,
        padding: const EdgeInsets.all(8.0),
        child: GridView.builder(
          itemCount: 81,
          gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
            crossAxisCount: 9,
            childAspectRatio: 1.0,
          ),
          itemBuilder: (context, index) {
            int row = index ~/ 9;
            int col = index % 9;
            return SudokuCell(
              value: sudokuBoard[row][col],
              notes: notes[row][col],
              isNoteMode: widget.isNoteMode,
              onChanged: (newValue) {
                setState(() {
                  if (widget.isNoteMode) {
                    // Add or remove notes based on the user's input
                    for (int note in newValue) {
                      if (notes[row][col].contains(note)) {
                        notes[row][col].remove(note);
                      } else {
                        notes[row][col].add(note);
                      }
                    }
                  } else {
                    if (newValue == [0]) {
                      // Clear the cell and show the previous notes
                      sudokuBoard[row][col] = 0;
                    } else {
                      // Set the answer and clear the notes
                      sudokuBoard[row][col] = newValue.first;
                      // notes[row][col].clear();
                    }
                  }
                });
              },
              onCleared: () {
                setState(() {
                  // When the cell is cleared, make sure the notes become visible again
                  sudokuBoard[row][col] = 0;
                });
              },
            );
          },
        ),
      ),
    );
  }
}

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
