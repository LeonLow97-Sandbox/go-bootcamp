# Liskov Substitution Principle (LSP)

The Liskov Substitution Principle states that objects of a superclasss should be replaceable with objects of a subclass without affecting the correctness of the program. This means that derived classes must adhere to the behavior expected by the base class.

## Importance

- Ensures that the program behaves as expected when substituting subclasses for their base classes.
- Promotes robust, maintainable code by adhering to consistent interfaces.

## Application of LSP in Code Example

- **Interface Definition**
  - The `Sized` interface defines methods for getting and setting width and height, which both `Rectangle` and `Square` are expected to implement.
- **Rectangle Struct**
  - The `Rectangle` struct correctly implements the `Sized` interface:
    - `GetWidth`, `GetHeight`: return the respective dimensions
    - `SetWidth`, `SetHeight`: set dimensions independently without altering each other.
- **Square Struct**
  - The `Square` struct inherits from `Rectangle` but violates LSP:
    - When `SetWidth` or `SetHeight` is called, it changes both dimensions, which is not the expected behavior for a rectangle.
    - This leads to unexpected results when substituting a `Square` for a `Rectangle` in functions that expect consistent behavior.
- **Alternative Approach (Square2)**
  - The `Square2` struct is designed to avoid LSP violation:
    - It implements its own size management without inheriting from `Rectangle`
    - By maintaining a single dimension (`size`), it adheres to LSP while still fulfilling the `Sized` interface.

## Key Takeaways

- **LSP Violation**: The original `Square` class demonstrates how inheritance can lead to violations of LSP, as it alters expected behavior be having interdependent width and height.
- **Composition Over Inheritance**: The approach taken with `Square2` illustrates the importance of composition. By not relying on inheritance from `Rectangle`, it maintains consistent behavior and avoids unexpected side effects.
- **Consistent Behavior**: For LSP to hold, all derived classes must behave in a manner consistent with the base class's contract, ensuring that substituting one for another does not alter program correctness.

## Summary

The code illustrates both the correct and incorrect applications of the Liskov Substitution Principle. By recognizing the violation in the `Square` struct and adopting a more appropriate design with `Square2`, it demonstrated an understanding of how to ensure consistency and correctness when implementing polymorphism. This principle is crucial for creating robust and maintainable software.
