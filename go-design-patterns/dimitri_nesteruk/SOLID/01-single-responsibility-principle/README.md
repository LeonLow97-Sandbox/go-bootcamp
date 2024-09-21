# Single Responsibility Principle (SRP)

- SRP states that a class or module should have one, and only one, reason to change. This means each class should have a specific focus or responsibility.

## Benefits

- Improved maintainability: Easier to modify or extend functionality
- Enhanced readability: Clearer purpose of each class or module
- Simplified testing: Easier to test specific functionalities without interference from other concerns

## Applications of SRP in Code Example

### `Journal` struct

- **Responsibility**: Managing journal `entries`
- **Methods**:
  - `AddEntry(text string)`: Adds a new entry.
  - `RemoveEntry(index int)`: Removes an entry by index.
  - `String() string`: Returns all entries as a formatted string.
- **No Persistence Logic**: The `Journal` does not handle file I/O operations, keeping its focus narrow.

### `Persistence` struct

- **Responsibility**: Handling the saving and loading of journal data.
- **Methods**:
  - `SaveToFile(j *Journal, filename string)`: Saves journal entries to a specified file.
- Separation from `Journal`: By isolating file operations, the `Persistence` struct adheres to SRP and allows for changes in persistence logic without affecting the `Journal`

### Summary

- The `Journal` and `Persistence` structs exemplify the Single Responsibility Principle by clearly stating their roles. This separation allows each component to evolve independently, facilitating easier maintenance and comprehension of the codebase. Always aim for clear responsibilities when designing classes or modules to adhere to SRP effectively.
