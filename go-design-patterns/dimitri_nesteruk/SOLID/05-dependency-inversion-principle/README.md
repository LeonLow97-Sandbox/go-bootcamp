# Dependency Inversion Principle (DIP)

The Dependency Inversion Principle states that high-level modules should not depend on low-level modules. Both should depend on abstractions (interfaces on Go).

## Importance

- Promotes a flexible and maintainable design.
- Reduces coupling between high-level and low-level modules, making it easier to change and extend the code.

## Analysis of Code Example

- **Abstraction**
  - The `RelationshipBrowser` interface abstracts the functionality required for retrieving relationships. This allows high-level code to depend on this abstraction rather than on a specific implementation.
- **Low-Level Module**
  - The `Relationships` struct implements the `RelationshipBrowser` interface, encapsulating the logic for managing relationships. This is the low-level module that provides concrete behavior.
- **High-Level Module**
  - The `Research` struct depends on the `RelationshipBrowser` interface rather than the `Relationships` struct directly. This ensures that the high-level logic does not depend on specific implementation of the low-level module, adhering to DIP.
- **Investigate Method**
  - The `Investigate` method in the `Research` struct uses the `FindAllChildrenOf` method from the `RelationshipBrowser` interface, demonstrating how high-level logic can operate independently of specific implementations.
- **Main Function**
  - In the `main` function, we create instances of `Person` and `Relationships`, and then pass the `Relationships` instance to Research. This showcases dependency injection, where dependencies are provided at runtime rather than hardcoded.

## Summary of Code Implementation

- Defining abstractions (interfaces) that high-level modules depend on.
- Keeping the high-level module (`Research`) decoupled from the low-level module (`Relationships`).
- Allowing for flexibility and maintainability by enabling the easy swapping of different implementations of `RelationshipBrowser`.

## Notes

- **Principle Overview**
  - High-level modules should not depend directly on low-level modules; both should depend on abstractions (interfaces).
  - This principle is fundamental for achieving a decoupled architecture.
- **Benefits**
  - Flexibility: You can change low-level implementations without affecting high-level code.
  - Testability: High-level modules can be tested in isolation by mocking dependencies.
  - Maintainability: Easier to manage and extend the codebase as it grows.
- **Implementation**
  - Define interfaces that encapsulate the behavior required by high-level modules.
  - Implement these interfaces in low-level modules, allowing high-level modules to depend on the interfaces instead of concrete implementations.
- **Examples in Code**:
  - Use interfaces (e.g., `RelationshipBrowser`) to abstract functionality.
  - Implement concrete types (e.g., `Relationships`) that fulfill these interfaces.
  - Inject dependencies (e.g., passing `Relationships` to `Research`) to decouple high-level and low-level code.
