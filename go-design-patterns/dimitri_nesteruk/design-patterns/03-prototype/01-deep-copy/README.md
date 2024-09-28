# Deep Copy

The deep copy pattern involves creating a new instance of an object that is a copy of an existing object, including all objects referenced by the original object. This ensures that changes made to the copied object do not affect the original object, which is crucial for maintaining data integrity in complex structures.

## Code Analysis

---

- `DeepCopy` method for `Address`
  - `DeepCopy` method for the `Address` struct creates a new `Address` instance and copies over the values from the original.
  - This ensures that if the copied `Address` is modified, it does not affect the original `Address`

```go
func (a *Address) DeepCopy() *Address {
    return &Address{
        a.StreetAddress,
        a.City,
        a.Country,
    }
}
```

---

- `DeepCopy` method for `Person`
  - The `DeepCopy` method for the `Person` struct first creates a shallow copy of the Person instance using `q := *p`. However, this only copies the pointer to the `Address`, which can lead to shared references.
    - What is a shallow copy? A shallow copy create s anew object that is a copy of the original object, but it does not create copies of the objects that the original object references. Instead, it copies the references to those objects. This means that if the original object contains references to mutable objects (like slices, maps, or pointers), both the original and the copied object will point to the same underlying data.
  - To ensure a true deep copy, it calls the `DeepCopy` method on the `Address` to get a new instance.
  - For the `Friends` slice, `copy(q.Friends, p.Friends)` is creates a new slice and copy the contents.

```go
func (p *Person) DeepCopy() *Person {
    q := *p
    q.Address = p.Address.DeepCopy()
    copy(q.Friends, p.Friends)
    return &q
}
```

---

- Main Function
  - An instance of `Person`, `john` is created.
  - A deep copy of `john` is made as `jane`. Changes made to `jane` (like modifying the name and address, and adding a friend) do not affect `john`.
  - The output clearly shows that `john` retains its original values, demonstrating that a true deep copy was made.

---

## Summary

The deep copy pattern is crucial for managing complex data structures where nested objects need to be modified independently. The approach in the code effectively demonstrates this concept through the `DeepCopy` methods, ensuring that all references are copied appropriately.
