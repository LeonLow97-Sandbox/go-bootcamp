# Builder Parameter

The builder pattern is used here to simplify the construction of complex objects (`email` in this case) by allowing the client code to specify only the required parameters step by step. Key aspects of the builder pattern include: - **Fluent Interface**: Each method (`From, To, Subject, Body`) returns `*EmailBuilder`, enabling method chaining (`builder.From().To().Subject().Body()`). - **Separation of Construction and Representation**: The construction process (`EmailBuilder`) is separate from the representation (`email` struct). - **Immutable Representation**: The `email` struct itself remains immutable during construction, avoiding inconsistent or incomplete states.

## Benefits

- **Clarity and Readability**: The fluent interface makes it clear how an `email` is being constructed and what components it consists of.
- **Flexibility**: The pattern allows adding optional parameters easily without cluttering the constructor, and it facilitates the creation of different configurations of the same object.
- **Encapsulation**: The `email` struct is encapsulated within the `EmailBuilder`, preventing direct modification and enforcing the construction process.
