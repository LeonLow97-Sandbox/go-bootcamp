# Go Design Patterns

- Design Patterns are typically OOP-based
- Go is not OOP
  - No inheritance
  - Weak encapsulation
- Permissive visibility/naming
- Use of OOP-ish terminology
  - _Hierarchy_ implies a set of related types: shared interface, embedding
  - _Properties_ are a combination of getter/setter methods

## SOLID Design Principles

|               Principle               | Description                                                                                                                                                                         |
| :-----------------------------------: | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Single Responsibility Principle (SRP) | A class should have only 1 reason to change, or in other words, it should have a single responsibility.                                                                             |
|      Open/Closed Principle (OCP)      | Software entities should be open for extension but closed for modification.                                                                                                         |
|  Liskov Substitution Principle (LSP)  | Subtypes must be substitutable for their base types without altering the correctness of the program.                                                                                |
| Interface Segregation Principle (ISP) | Clients should not be forced to depend on interfaces they do not use.                                                                                                               |
| Dependency Inversion Principle (DIP)  | High-level modules should not depend on low-level modules; both should depend on abstractions. Abstractions should not depend on details, but details should depend on abstractions |