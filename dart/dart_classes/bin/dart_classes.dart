import 'user_class.dart';

void main(List<String> arguments) {
  // Cascade operator chains together mulyiple assignments
  // final user = User()
  //   ..id = 2
  //   ..name = "Tremor";

  final anonymousUser = User.anonymous();

  // factory constructor
  final json = {'id': 10, 'name': "Tremor2"};
  final tremor2 = User.fromJson(json);

  print(tremor2);
  print(anonymousUser);
}

