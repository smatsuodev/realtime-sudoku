import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:frontend/domain/entity/room.dart';
import 'package:frontend/domain/repository/room_repository.dart';
import 'package:frontend/presentation/component/room_preview_card.dart';

class RoomsListPage extends ConsumerStatefulWidget {
  const RoomsListPage({super.key});

  @override
  ConsumerState<ConsumerStatefulWidget> createState() => _RoomsListPageState();
}

class _RoomsListPageState extends ConsumerState<RoomsListPage> {
  late Future<List<Room>> _rooms;
  late final _roomRepository = ref.read(roomRepositoryProvider);

  @override
  void initState() {
    _rooms = _fetchRooms();
    super.initState();
  }

  Future<List<Room>> _fetchRooms() async {
    return await _roomRepository.getRooms();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('ルーム一覧'),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          setState(() {
            _rooms = _fetchRooms();
          });
        },
        tooltip: 'ルームを新規作成',
        child: const Icon(Icons.add),
      ),
      body: Center(
        child: FutureBuilder(
            future: _rooms,
            builder: (context, snapshot) {
              if (snapshot.hasError) {
                return const Text('エラーが発生しました');
              }
              if (!snapshot.hasData ||
                  snapshot.connectionState != ConnectionState.done) {
                return const CircularProgressIndicator();
              }

              final rooms = snapshot.data as List<Room>;

              return SizedBox(
                width: 1000,
                child: GridView(
                  padding: const EdgeInsets.only(
                      left: 32, right: 32, top: 16, bottom: 16),
                  gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                    crossAxisCount: 4,
                  ),
                  children: rooms
                      .map((room) => RoomPreviewCard(
                            room: room,
                            onTap: () {
                              print(room.name);
                            },
                          ))
                      .toList(),
                ),
              );
            }),
      ),
    );
  }
}
