# Adapter

- Adapter is a construct which adapts an existing interface X to conform to the required interface Y.
- ChatGPT: The Adapter Pattern is a structural design pattern that allows objects with incompatible interfaces to work together. It acts as a bridge between 2 incompatible interface by converting the interface of a class into another interface that the client expects. This is particularly useful when you want to reuse existing code that doesn't match the required interface.

# Key Concepts

- **Target Interface**: The interface that the client expects
- **Adaptee**: The existing interface that needs adaptation
- **Adapter**: The class that implements the target interface and holds an instance of the adaptee, converting calls to the adaptee's interface as needed.

## How the code implements the Adapter Pattern

The Adapter Pattern is implemented to adapt a `VectorImage` to a `RasterImage`

---

1. **Adaptee**
   - `VectorImage` represents the existing interface that the code uses to define shapes using lines. It has a specific structure (lines) that needs to be converted to a different form (points).
2. **Target Interface**
   - `RasterImage` is the interface that the `DrawPoints` function expects. It defines method `GetPoints()` that returns a list of `Point` objects.
3. **Adapter**
   - `vectorToRasterAdapter` acts as the adapter. It implements the `RasterImage` interface, specifically the `GetPoints()` method. Internally, it converts the `Line` data from the `VectorImage` into `Point` data that can be used by `DrawPoints`

---

# Summary

- Implementing an Adapter is easy
- Determine the API you gave and the APi you need
- Create a component which aggregates (has a pointer to ...) the adaptee
- Intermediate representations can pile up: use caching and other optimizations
- The code example demonstrates the Adapter Pattern by providing a mechanism to adapt `VectorImage` (which consists of lines) into a `RasterImage` (which consists of points). The `vectorToRasterAdapter` effectively allows `DrawPoints` to work with `VectorImage` without modifying the original vector structure, thereby promoting code reusability and separation of concerns.
