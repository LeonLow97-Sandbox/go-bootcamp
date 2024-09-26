# Prototype Factory

A prototype factory is a design pattern that creates objects based on a template or prototype. Instead of creating new instances from scratch, it uses a predefined prototype to produce new instances. This is particularly useful when the cost of creating an object is more expensive than cloning an existing instance, or when the object configuration is complex.

## Key Characteristics of Prototype Factory

1. **Template/Prototype**: It uses existing objects as prototypes.
2. **Cloning**: New objects are created by copying the prototype rather than by defining their properties from scratch.
3. **Encapsulation of Creation Logic**: The creation logic can be centralized, allowing for easier maintenance and extension.

## Code Analysis

- **Predefined Prototypes**: Defined 2 types of employees, `Developer` and `Manager`, with predefined characteristics (like position and income).
- **Creation Logic**: The `NewEmployee` function encapsulates the logic for creating employees based on their roles. It returns a new instance of the `Employee` struct with predefined values for `Position` and `AnnualIncome`.
- **Cloning Behavior**: Although not a classic clone in the sense of object cloning, you are creating new `Employee` instances based on predefined roles. This allows for consistency in object creation and makes it easy to manage employee types.
