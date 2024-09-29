# Bridge

- A mechanism that decouples an interface (hierarchy) from an implementation (hierarchy).
- The bridge design pattern is a structural pattern that **separates an abstraction from its implementation**, allowing both to evolve independently. It is particularly useful when you want to decouple an interface from its implementation, providing flexibility in how the implementation can be changed without affecting the client code.

# Key Component of the Bridge Pattern

- **Abstraction**: This is the high-level interface that defines the operations that can be performed. It holds a reference to the implementation interface.
- **Implementor**: This is the interface that defines the methods for the actual implementations. Different concrete classes can implement this interface.
- **Concrete Implementation**: These are the classes that implement the implementor interface.
- **Refined Abstraction**: This extends the abstraction and may add additional behaviors but still relies on the implementor.

# How the Code Implements the Bridge Pattern

- **Abstraction**: The `Circle` struct acts as the abstraction. It has methods like `Draw` and `Resize`, which operate on the `Renderer` interface.
- **Implementor**: The `Renderer` interface defines the method `RenderCircle(radius float32)`. This represents the oeprations that the `Circle` abstraction can call.
- **Concrete Implementations**: The `VectorRenderer` and `RasterRenderer` structs implement the `Renderer` interface. Each provides its own version of the `RenderCircle` method.
  - `VectorRenderer` draws a circle using vector graphics.
  - `RasterRenderer` draws a circle using raster graphics.
- **Refined Abstraction**: The `Circle` struct serves as the refined abstraction, which can work with any implementation of the `Renderer` interface, allowing it to be drawn in different ways (either raster or vector)

# Benefits of the Approach

- **Decoupling**: The `Circle` class is decoupled from the specific rendering implementation. You can add new renderers (like `RasterRenderer` or `VectorRenderer`) without modifying the `Circle` class.
- **Flexibility**: Can easily switch the rendering strategy at runtime by changing the renderer passed to the `Circle`. For example, you can create a circle that uses `RasterRenderer` instead of `VectorRenderer` simply by changing the initialization.
- **Scalability**: If you decide to add new shapes (like squares or triangles), you can do so without altering the existing code structure significantly.

# Summary

- Decouple from abstraction from implementation
- Both can exist as hierarchies
- A stronger form of encapsulation
