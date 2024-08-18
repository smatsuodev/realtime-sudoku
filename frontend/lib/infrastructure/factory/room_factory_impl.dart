import 'package:frontend/domain/entity/room.dart';
import 'package:frontend/domain/entity/user.dart';
import 'package:frontend/domain/factory/room_factory.dart';
import 'package:frontend/domain/factory/user_factory.dart';
import 'package:frontend/infrastructure/model/room.dart';

class RoomFactoryImpl implements RoomFactory {
  final UserFactory _userFactory;

  RoomFactoryImpl({
    required UserFactory userFactory,
  }) : _userFactory = userFactory;

  @override
  Room create({
    required int id,
    required String name,
    required List<User> players,
  }) {
    return Room(
      id: id,
      name: name,
      players: players,
    );
  }

  @override
  Room createFromModel(RoomModel model) {
    return Room(
      id: model.id,
      name: model.name,
      players: model.players.map(_userFactory.createFromModel).toList(),
    );
  }
}
