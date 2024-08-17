import 'package:flutter/material.dart';
import 'package:frontend/domain/entity/room.dart';

class RoomPreviewCard extends StatelessWidget {
  final Room room;
  final Function() onTap;

  const RoomPreviewCard({
    super.key,
    required this.room,
    required this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    return Card(
      child: InkWell(
        customBorder: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(8),
        ),
        onTap: onTap,
        child: Padding(
          padding: const EdgeInsets.all(8.0),
          child: Column(
            children: [
              const Expanded(
                // TODO: 部屋のボードの画像を表示する
                child: Icon(Icons.no_photography_outlined),
              ),
              Text(room.name),
            ],
          ),
        ),
      ),
    );
  }
}
