import 'package:flutter/material.dart';
import 'package:frontend/domain/entity/room.dart';

class RoomsListPage extends StatelessWidget {
  final List<Room> _rooms;

  const RoomsListPage({
    super.key,
    required List<Room> rooms,
  }) : _rooms = rooms;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('ルーム一覧'),
      ),
      body: Center(
        child: SizedBox(
          width: 1000,
          child: GridView(
            padding:
                const EdgeInsets.only(left: 32, right: 32, top: 16, bottom: 16),
            gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
              crossAxisCount: 4,
            ),
            children: _rooms
                .map((room) => RoomCard(
                      roomName: room.name,
                      onTap: () {
                        print(room.name);
                      },
                    ))
                .toList(),
          ),
        ),
      ),
    );
  }
}

class RoomCard extends StatelessWidget {
  final String roomName;
  final Function() onTap;

  const RoomCard({
    super.key,
    required this.roomName,
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
              Text(roomName),
            ],
          ),
        ),
      ),
    );
  }
}
