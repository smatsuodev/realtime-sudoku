import 'package:frontend/domain/entity/room.dart';
import 'package:frontend/domain/repository/room_repository.dart';

class MockRoomRepositoryImpl implements RoomRepository {
  final bool shouldDelay;
  final int delaySeconds;
  final int numberOfRooms;

  late final _rooms = List.generate(
    numberOfRooms,
    (index) => Room(
      id: index,
      name: 'Room $index',
      players: [],
    ),
  );

  MockRoomRepositoryImpl({
    this.shouldDelay = false,
    this.delaySeconds = 1,
    this.numberOfRooms = 4,
  });

  @override
  Future<void> createRoom(String name) {
    return delayed<void>(() {
      _rooms.add(Room(id: _rooms.length, name: name, players: []));
    });
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

  Future<T> delayed<T>(T Function() f) async {
    if (shouldDelay) {
      return Future.delayed(
        Duration(seconds: delaySeconds),
        () => f(),
      );
    }
    return f();
  }
}
