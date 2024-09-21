# Open-Closed Principle (OCP)

The Open/Closed Principle states that software entities (like classes, modules, and functions) should be open for extension but closed for modification. This means you can add new functionality without altering existing code.

## Importance

- Promotes maintainability and stability in code.
- Reduces the risk of introducing bugs when adding new features as you already testing the existing functions.

## Application of OCP in the code example

- **Specification Pattern**:
  - The implementation uses the Specification pattern, which allows for easy extension of filtering criteria.
  - New specifications (like `ColorSpecification`, `SizeSpecification`, and `AndSpecification`) can be added without changing the existing filtering logic in the `BetterFilter`
- **Separation of Concerns**
  - The `BetterFilter` struct encapsulate filtering behavior. They do not need to change when you introduce new filtering criteria, adhering to the OCP.
- **Adding New Specification**
  - To add a new specification (e.g., `ShapeSpecification`), you simply create a new struct that implements the `Specification` interface. This does not require any modification to existing classes or methods.
  ```go
  type ShapeSpecification struct {
      shape string
  }
  ```
- **Use of Interfaces**
  - The `Specification` interface provides a contract that different specification structs implement. This abstraction allows `BetterFilter` to work with any new specifications seamlessly.
  - As a result, the `Filter` method in `BetterFilter` remain unchanged, fulfilling the requirement of being closed for modification.
- **Flexible Filtering**
  - The design allows you to combine specifications (e.g., using `AndSpecification`) to create more complex filters without modifying existing code. This enhances the system's capability while adhering to OCP.

## Summary

- The code example implements the Open/Closed Principle by utilizing the Specification pattern, allowing for easy extension of filtering the criteria without modifying existing logic.
- New specifications can be added by simply implementing the `Specification` interface, maintaining the stability of the existing codebase.
- The separation of concerns between filtering logic and specifications ensure that the code is adaptable and maintainable, embodying the principles of OCP.
