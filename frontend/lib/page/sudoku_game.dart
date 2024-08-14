class SudokuApp extends StatefulWidget {

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
