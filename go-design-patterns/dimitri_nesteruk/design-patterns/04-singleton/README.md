# Singleton

- For some components, it only makes sense to have one call in the system
  - Database repository
  - Object factory
- E.g., the construction call is expensive
  - We only do it once
  - We give everyone the same instance
- Want to prevent anyone creating additional copies
- Need to take care of lazy instantiation, "lazy loading"
- **Singleton**: A component is instantiated only once.

# Summary

- Lazy one-time initialization using `sync.Once`, handles thread safety
- Adhere to DIP: depend on interfaces (`Database`), not concrete types (`GetSingletonDatabase`).
- Singleton is not scary...
  - Be careful of depending directly on singleton only, create an interface
