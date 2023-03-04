// Class
class User {
  // Generative
  // generative un-named constructor
  const User({this.id = 0, this.name = "Anonymous"});

  // generative Named constructor
  // forwarding=> calling the main constructor frm the named constr... instead of
  // repeating assigning the var inside curly braces
  // User.anonymous() : this(0, "Anonymous");
  const User.anonymous() : this();

  // factory constructor
  // returns existing instances of a class ...while generatives create new instances

  factory User.frank() {
    return User(id: 200, name: "Frank");
  }

  // Example. 2
  // factory constr allows making some operations(validation, modifications) 
  //before returning the data
  factory User.fromJson(Map<String, Object> json) {
    final userId = json['id'] as int;
    final userName = json['name'] as String;

    return User(id: userId, name: userName);
  }

  final int id;
  final String name;

  // how the object will print out => print(user)
  @override
  String toString() {
    return "User{id: $id, name: $name}";
  }
}
