import 'package:frontend/domain/entity/room.dart';
import 'package:frontend/domain/repository/room_repository.dart';

class MockRoomRepositoryImpl implements RoomRepository {
  final _rooms = List.generate(
    16,
    (index) => Room(
      id: index,
      name: 'Room $index',
      players: [],
    ),
  );

  final bool shouldDelay;
  final int delaySeconds;

  MockRoomRepositoryImpl({
    this.shouldDelay = false,
    this.delaySeconds = 1,
  });

  @override
  Future<void> createRoom(String name) {
    // TODO: implement createRoom
    throw UnimplementedError();
  }

  @override
  Future<void> deleteRoom(int roomId) {
    // TODO: implement deleteRoom
    throw UnimplementedError();
  }

  @override
  Future<List<Room>> getRooms() async {
    if (shouldDelay) {
      return Future.delayed(
        Duration(seconds: delaySeconds),
        () => _rooms,
      );
    }

    return _rooms;
  }
}
