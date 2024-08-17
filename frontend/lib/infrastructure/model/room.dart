import 'package:frontend/infrastructure/model/user.dart';
import 'package:frontend/gen/sudoku/entities/v1/room.pb.dart' as pb;
import 'package:frontend/gen/sudoku/entities/v1/user.pb.dart' as pb;

class RoomModel {
  final int id;
  final String name;
  final List<UserModel> players;

  RoomModel({
    required this.id,
    required this.name,
    required this.players,
  });

  factory RoomModel.fromProto(pb.Room room) {
    return RoomModel(
      id: room.id,
      name: room.name,
      players: room.players
          .map((pb.User user) => UserModel.fromProto(user))
          .toList(),
    );
  }
}
