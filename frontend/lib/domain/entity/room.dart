import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:frontend/domain/entity/user.dart';

part 'room.freezed.dart';

@freezed
class Room with _$Room {
  const factory Room({
    required int id,
    required String name,
    required List<User> players,
  }) = _Room;
}
