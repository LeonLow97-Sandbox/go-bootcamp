# Singleton

The singleton pattern ensures that a class has only 1 instance and provides a global point of access to it. This is useful when exactly one object is needed to coordinate actions across the system, like managing a database connection or configuration settings.

# How Singleton was implemented in the code

1. **Private Constructor**: The `singletonDatabase` struct is designed to be instantiated only within the package. This prevents external code from creating new instances directly.
2. **Once Initialization**: The `sync.Once` type is used to guarantee that the initialization of the singleton instance occurs only once. The `once.Do` method takes a function, which is executed only once, ensuring thread safety.
3. **Global Access**: The `GetSingletonDatabase` function provides access to the singleton instance. If it hasn't been created yet, it reads the data from `capitals.txt` and initializes the `singletonDatabase`.
4. **Data Storage**: The actual data (capitals and their populations) is stored in a map within the singleton instance, which can be accessed via the `GetPopulation` method.

# Dependency Inversion Principle (DIP)

The dependency inversion principle (DIP) states that high-level modules should not depend on low-level modules; both should depend on abstractions. This promotes loose coupling and enhances testability.

## How did the code incorporate DIP?

1. **Database Interface**: The `Database` interface defines a method `GetPopulation(name string) int`. Both the `singletonDatabase` and `DummyDatabase` structs implement this interface. This allows for polymorphism, where you can use any type that satisfies the `Database` interface without knowing the specifics of its implementation
2. **High-Level Functionality**: The function `GetTotalPopulationEx` takes a `Database` interface as a parameter, allowing it to work with any database implementation. This decouples the function from the concrete implementation of the database, facilitating easier testing and flexibility.
3. **Unit Testing**: By using the `DummyDatabase`, which provides hardcoded data, you can test functions like `GetTotalPopulationEx` without needing to access an actual database. This ensures that your unit tests are fast and reliable, as they do not depend on external resources.
