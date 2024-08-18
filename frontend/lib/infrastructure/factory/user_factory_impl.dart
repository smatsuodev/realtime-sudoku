import 'package:frontend/domain/entity/user.dart';
import 'package:frontend/domain/factory/user_factory.dart';
import 'package:frontend/infrastructure/model/user.dart';

class UserFactoryImpl implements UserFactory {
  @override
  User create({
    required int id,
    required String name,
  }) {
    return User(
      id: id,
      name: name,
    );
  }

  @override
  User createFromModel(UserModel model) {
    return User(
      id: model.id,
      name: model.name,
    );
  }
}
