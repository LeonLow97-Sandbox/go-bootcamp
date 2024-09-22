# Builder Pattern

The builder pattern is a creational design pattern that allows for the step by step creation of complex objects. It separates the construction of an object from its representation, enabling the same construction process to create different representations.

## Code Analysis

- `HtmlElement` struct
  - Represents a single HTMl element with:
    - `name`: the tag name (e.g., "div", "p")
    - `text`: the inner text content
    - `elements`: a slice of child `HtmlElement`.
  - Contains methods for generating a string representation of the HTML.
- `String` Method
  - The `String()` method generates a formatted string of the HTML element.
  - Uses recursion to handle nested elements, with indentation for better readability.
- `HtmlBuilder` struct:
  - Serves as the builder for creating HTMl elements
  - Holds:
    - `rootName`: the name of the root HTML element
    - `root`: An `HtmlElement` representing the root of the structure
- Builder Methods
  - `NewHtmlBuilder(rootName string)`: initializes a new builder with a root element
  - `String()`: returns the string representation of the entire HTMl structure.
  - `AddChild(childName, childText string)` adds a child element to the root
  - `AddChildFluent(childName, childText string)`: adds a child and returns the builder for chaining calls.
- Usage
  - The main function demonstrates how to use the builder pattern:
    1. Create a simple HTML strings manually with a `strings.Builder`
    2. Use the `HtmlBuilder` to construct an HTML structure in a readable and maintainable way
    3. Utilize method chaining with `AddChildFluent()` for a fluent interface
- Advantages:
  - **Readability**: Clear and organized way to build complex structures
  - **Maintainability**: Easy to modify and extend without changing the overall structure.
  - **Separation of Concerns**: focuses on how to build the object, not on how it looks.
