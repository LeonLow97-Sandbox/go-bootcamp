# Interface Segregation Principle (ISP)

The interface segregation principle states that no client should be forced to depend on methods it does not use. Instead of having a large interface, it's better to have smaller, specific ones.

## Importance

- Promotes more modular and maintainable code.
- Reduces the impact of changes since clients depend only on the interfaces they need.

## Analysis of Code Example

- **Machine Interface**
  - The `Machine` interface with `Print`, `Fax`, and `Scan` methods. While this can be useful for a multifunction device, it can lead to issues with classes that do not support all functionalities (like `OldFashionedPrinter`).
- OldFashionedPrinter
  - The `OldFashionedPrinter` implements the `Machine` interface but provides panic methods for `Fax` and `Scan`. This violates ISP because clients using this printer are forced to handle unsupported operations.
- ISP Implementation:
  - Addressed ISP by creating smaller interfaces:
    - `Printer`: Contains only the `Print` method
    - `Scanner`: Contains only the `Scan` method
- MyPrinter and Photocopier
  - `MyPrinter` implements the `Printer` interface, and `Photocopier` implements both `Printer` and `Scanner`. This means clients can use them without worrying about unused methods, adhering to ISP.
- MultiFunctionMachine
  - The `MultiFunctionMachine` aggregated `Printer` and `Scanner`, allowing it to use only the methods it needs, thus adhering to ISP. This composition allows you to build multifunction devices without forcing all devices to implement all methods.

## Summary

- The code example correctly follow the ISP by creating smaller, more focused interfaces (`Printer` and `Scanner`), allowing clients to depend only on the methods they need.
- The `OldFashionedPrinter` example highlights the importance of not forcing clients to implement methods that are not applicable to them, which is a key ISP violation.
- By using composition in the `MultiFunctionMachine`, you create a flexible design that promotes modularity and maintainability.
- By applying ISP, you enhance the flexibility of your code, making it easier to maintain and extend in the future.
