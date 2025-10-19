# Solid Principles in Go

## Summary

- **Single Responsibility Principle (SRP)**:
  - A type should only have one reason to change.
  - _Separation of concerns_: different types/packages handling different, independent tasks/problems.
- **Open-Closed Principle (OCP)**:
  - Types should be open for extension but closed for modification.
- **Liskov Substitution Principle (LSP)**:
  - You should be able to substitute an embedding type in place of its embedded part.
  - ❌ Not used in Golang's context.
- **Interface Segregation Principle (ISP)**:
  - Don't put too much into an interface; split into separate interfaces.
- **Dependency Inversion Principle (DIP)**:
  - High-level modules should not depend upon low-level ones; use abstractions.
