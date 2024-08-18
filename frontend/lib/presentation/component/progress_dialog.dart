import 'package:flutter/material.dart';

void showProgressDialog(BuildContext context) {
  showGeneralDialog(
    context: context,
    barrierDismissible: false,
    transitionDuration: Duration.zero,
    barrierColor: Colors.black.withOpacity(0.5),
    pageBuilder: (
      BuildContext context,
      Animation animation,
      Animation secondaryAnimation,
    ) {
      return const Center(
        child: CircularProgressIndicator(),
      );
    },
  );
}

void hideProgressDialog(BuildContext context) {
  Navigator.of(context, rootNavigator: true).pop();
}
