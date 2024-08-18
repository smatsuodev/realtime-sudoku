import 'package:frontend/gen/sudoku/entities/v1/user.pb.dart' as pb;

class UserModel {
  final int id;
  final String name;

  UserModel({
    required this.id,
    required this.name,
  });

  factory UserModel.fromProto(pb.User user) {
    return UserModel(
      id: user.id,
      name: user.name,
    );
  }
}
