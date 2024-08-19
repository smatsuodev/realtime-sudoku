import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

class SudokuCell extends StatelessWidget {
  final int value;
  final Function(int newValue) onChanged;

  SudokuCell({
    super.key,
    required this.value,
    required this.onChanged,
  });

  KeyEventResult _handleKeyPress(FocusNode node, KeyEvent event) {
    if (event is KeyDownEvent) {
      final key = event.logicalKey.keyId - LogicalKeyboardKey.digit0.keyId;
      if (0 <= key && key <= 9) {
        this.onChanged(key);
        return KeyEventResult.handled;
      }
    }
    return KeyEventResult.ignored;
  }

  @override
  Widget build(BuildContext context) {
    return Focus(
      onKeyEvent: _handleKeyPress,
      child: Builder(builder: (context) {
        final focusNode = Focus.of(context);
        final hasFocus = focusNode.hasFocus;
        return GestureDetector(
          onTap: () {
            if (hasFocus) {
              focusNode.unfocus();
            } else {
              focusNode.requestFocus();
            }
          },
          child: Container(
            width: 100,
            height: 100,
            decoration: BoxDecoration(
              border: Border.all(color: Colors.black),
              color: hasFocus ? Colors.yellow : Colors.white,
            ),
            child: Center(
              child: Text(
                value.toString(),
                textAlign: TextAlign.center,
                style: TextStyle(fontSize: 24),
              ),
            ),
          ),
        );
      }),
    );
  }
}
