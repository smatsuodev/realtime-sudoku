import 'package:frontend/domain/entity/user.dart';
import 'package:frontend/infrastructure/model/user.dart';

abstract class UserFactory {
  User create({
    required int id,
    required String name,
  });

  User createFromModel(UserModel model);
}
