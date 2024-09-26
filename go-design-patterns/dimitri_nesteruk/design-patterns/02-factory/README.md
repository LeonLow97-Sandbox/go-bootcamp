# Factory

- Object creation logic becomes too convoluted
- Struct has too many fields, need to initialize all correctly
- Wholesale object creation (non-piecewise, unlike Builder) can be outsourced to
  - A separate function (Factory Function AKA Constructor)
  - They may exist in a separate struct (Factory)

## Summary

- A _factory function_ (aka constructor) is a helper function for making struct instances.
- A factory is any entity that can take care of object creation.
- Can be a function or a dedicated struct.
