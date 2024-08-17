import 'package:frontend/domain/entity/room.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'room_repository.g.dart';

@riverpod
RoomRepository roomRepository(RoomRepositoryRef ref) {
  throw UnimplementedError();
}

abstract class RoomRepository {
  Future<void> createRoom(String name);
  Future<void> deleteRoom(int roomId);
  Future<List<Room>> getRooms();
}
