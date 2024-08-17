import 'package:flutter/material.dart';
import 'package:frontend/domain/entity/room.dart';
import 'package:frontend/presentation/component/create_room_button.dart';
import 'package:frontend/presentation/component/room_preview_card.dart';

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
      floatingActionButton: const CreateRoomButton(),
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
                .map((room) => RoomPreviewCard(
                      room: room,
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
