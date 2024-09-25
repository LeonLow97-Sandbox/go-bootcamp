# Interface Factory

- An interface factory is a design pattern that uses an interface to create different types of objects based on specific conditions. The factory function **returns an instance of an interface rather than a concrete type**, allowing for flexibility and abstraction. This approach promotes encapsulation and reduces dependencies on specific implementations.
  - Encapsulation in OOP restricts direct access to some of an object's components, e.g., in our code example, users can call `SayHello()` without needing to directly manipulate or know about the `name` and `age` fields.

## Code Analysis

- The `Person` interface defines a single method, `SayHello()`. Any type that implements this method can be treated as a `Person`.
- There are two concrete types: `person` and `tiredPerson`. Each type has fields `name` and `age` and implements the `SayHello()` method.
- Factory Function:
  - The `NewPerson` function is the core of the interface factory. It takes parameters (name and age) and returns a `Person` interface:
    - If the age is over 100, it returns an instance of `tiredPerson`
    - Otherwise, it returns an instance of `person`

---

```go
p := NewPerson("James", 34)
p.SayHello()

tp := NewPerson("James", 200)
tp.SayHello()
```

## Can create instances of `Person` without directly dealing with the concrete types. This encapsulation prevents access to the underlying fields (`name` and `age`) since the concrete types are private (lowercase).

## Summary

- **Flexibility**: The factory can return different implementations based on the input parameters, allowing for varied behavior without changing the calling code.
- **Encapsulation**: Users interact with the `Person` interface, not with the concrete implementations, promoting loose coupling.
- **Abstraction**: The concrete details of how `SayHello` is implemented are hidden, which simplifies the interaction with the `Person` type.
- This approach is beneficial in larger applications where the types and behaviors may evolve independently, enhancing maintainability and scalability.
