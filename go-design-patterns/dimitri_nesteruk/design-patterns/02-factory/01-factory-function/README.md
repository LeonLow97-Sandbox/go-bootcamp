# Factory Function

The Factory Function pattern is a creational design pattern that uses factory methods to create objects without specifying the exact class of object that will be created. Instead of directly using a constructor (which in Go is typically just creating a struct instance directly), the factory function handles the creation and initialization logic. This allows for more control over the creation process, such as performing validations, pre-processing, or selecting from different implementation options.

## Code Analysis

- `NewPerson` function: This function acts as a factory function that encapsulates the logic for creating and initializing a `Person` object. It takes `name` and `age` parameters, which are required to create a `Person`. Inside this function:
  - It can perform validation checks.
  - It initializes the `Person` struct with the provided `name` and `age`, and sets `EyeCount` to a default value of 2.
  - It returns a pointer (`*Person`) to the newly created `Person` object.

## Factory Function Pattern in Code

- **Encapsulation**: The factory function `NewPerson` encapsulates the logic for creating and initializing `Person` objects. It hides the internal implementation details of how a `Person` object is created.
- **Flexibility**: It provides flexibility to change the initialization logic of `Person` objects in a centralized place (inside `NewPerson`), without affecting the calling code (`main()` in this case).
- **Control**: Allows for complex validation or logic before creating the object.

## Summary

- The provided code fulfills the Factory Function pattern because it abstracts the creation of `Person` objects into a separate function (`NewPerson`). This function handles the creation process, initializes the fields of the `Person` struct, and returns a pointer to the newly created object. This pattern promotes encapsulation, flexibility, and control over object creation, adhering to the principles of good software design.
