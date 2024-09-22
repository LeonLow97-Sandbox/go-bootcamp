# Builder Facet Pattern

The Builder Facet Pattern is a variation of the Builder Pattern that allows for the construction of an object using multiple facets (or aspects). Each facet focuses on a specific area of the object, enabling a clear and organized way to build complex objects with multiple attributes.

## Key Concepts

1. `Person` struct

- Represents an individual with:
  - Address information: `StreetAddress`, `Postcode`, `City`
  - Job Information: `CompanyName`, `Position`, `AnnualIncome`

2. `PersonBuilder` struct

- Main builder that aggregates the construction process
- Contains a pointer to a `Person` instance and provides methods to access specific builders for address and job facets.

3. Facet Builders

- `PersonAddressBuilder` manages the address-related attributes.
- `PersonJobBuilder` manages job-related attributes.
- Both facets are designed to allow chaining of methods, improving readability and ease of use.

4. Method Definitions

- The `Lives()` and `Works()` methods return instances of their respective builders, allowing users to switch contexts between building an address and a job.
- Each facet builder has methods that set specific attributes and return the builder for fluent interface chaining.

5. `Build` Method

- The `Build()` method in `PersonBuilder` returns the constructed `Person` object.

## Advantages

- **Separation of Concerns**: Each aspect of the object is built independently, making the code cleaner and easier to maintain
- **Fluent Interface**: Chaining methods enhances readability and provides a natural way to construct objects
- **Flexibility**: Allows users to focus on specific facets of the object without being overwhelmed by unrelated properties.a

## Summary

The Builder Facet Pattern is an effective way to create complex objects with multiple aspects. By providing distinct builders for different facets, this pattern provides clean and organized code, enhances readability, and allows for easy modifications. It is particularly useful in scenarios where objects have numerous optional attributes that can be grouped logically.
