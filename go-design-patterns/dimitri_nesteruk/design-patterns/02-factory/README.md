# Factory

- Object creation logic becomes too convoluted
- Struct has too many fields, need to initialize all correctly
- Wholesale object creation (non-piecewise, unlike Builder) can be outsourced to
  - A separate function (Factory Function AKA Constructor)
  - They may exist in a separate struct (Factory)
