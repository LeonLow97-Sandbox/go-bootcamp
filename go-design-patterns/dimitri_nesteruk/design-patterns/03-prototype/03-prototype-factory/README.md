# Prototype Factory

## Benefits of this approach

1. **Automatic Handling of Nested Structures**: Serialization allows for automatic copying of all nested fields, eliminating the need for manual copying and reducing the risk of errors.
1. **Prototype Pattern**: The use of prototypes (`mainOffice` and `auxOffice`) facilitates the creation of new employees with default values that can be easily customized, following the prototype design pattern.
1. **Decoupling**: The deep copy process is separated from the structure itself, leading to cleaner and more maintainable code.
1. **Independence**: Each new `Employee` instance is independent of its prototype, ensuring that modifications do not affect the original prototypes.
