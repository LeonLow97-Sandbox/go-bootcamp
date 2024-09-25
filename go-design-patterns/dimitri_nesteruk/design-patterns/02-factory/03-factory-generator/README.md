# Factory Generator

The Factory Generator pattern is a design pattern that provides a way to create objects (in this case, `Employee` instances) without exposing the instantiation logic to the client. Instead of using a direct constructor for creating objects, the pattern involves using functions (or methods) that encapsulate the logic for creating instances, allowing for flexibility and abstraction.

## Code Analysis

- In the provided Go code, 2 variations of the factory generator pattern are implemented
  - Functional Style
  - Struct-based style

---

1. **Functional Factory**: `NewEmployeeFactory`

```go
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
    return func(name string) *Employee {
        return &Employee{name, position, annualIncome}
    }
}
```

- **Function**: This function returns another function that can create an `Employee` instance.
- **Parameters**: It takes a `position` and `annualIncome`, which are used to set these properties for each new employee.
- **Closure**: The inner function captures the `position` and `annualIncome, allowing you to create employees with these fixed attributes without having to specify them each time.
- **Usage**: When you call `NewEmployeeFactory`, it returns a function that takes a `name` and creates an `Employee`, allowing for flexible instantiation.

2. **Struct-Based Factory**: `EmployeeFactory`

```go
type EmployeeFactory struct {
    Position     string
    AnnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
    return &Employee{name, f.Position, f.AnnualIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
    return &EmployeeFactory{position, annualIncome}
}
```

- **Struct**: This version uses a struct (`EmployeeFactory`) to hold the position and annual income, making it easy to create multiple employees with the same attributes.
- **Method**: The `Create` method takes a `name` and constructs an `Employee` using the fields of the struct.
- **Manipulation**: You can create an instance of `EmployeeFactory` and modify its attributes (like `AnnualIncome`) before creating employees, providing more flexibility compared to the functional approach.

---

## Summary

- Both styles of the factory generator pattern in this code offer a way to encapsulate the creation logic of `Employee` instances, allowing you to create employees with predefined attributes. The functional style provides simplicity and closure-based encapsulation, while the struct-based style offers more flexibility with the ability to modify attributes after instantiation. Understanding these 2 implementations will help you utilize the factory pattern effectively in your Go programs.
