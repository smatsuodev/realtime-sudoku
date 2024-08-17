import 'package:frontend/domain/entity/room.dart';
import 'package:frontend/domain/entity/user.dart';
import 'package:frontend/infrastructure/model/room.dart';

abstract class RoomFactory {
  Room create({
    required int id,
    required String name,
    required List<User> players,
  });

  Room createFromModel(RoomModel model);
}
