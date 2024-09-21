## Quiz Questions

1. Whenever you pass an integer, float, string or struct into a function, what does Go do with that argument?
- It creates a copy of each argument, and these copies are used inside of the function.

2. What is the `&` operator used for?
- Turning a value into a pointer

3. When you see a `*` operator in front of a pointer, what will it turn the pointer into?
- A value

4. In the 'changeLatitude' function, what is the *location in the receiver list (after the word 'func') communicating to us?
```
func (lo *location) changeLatitude() {
 (*lo).latitude = 41.0
}
```
- It specifies the type of the receiver that the function expects.
- Normally, a `*` means that this function expects to be called with a pointer to a location. In this case, its specifying a type and is not being used as an  operator.

5. When we create a slice, Go will automatically create which 2 data structures?
- An array and a structure that records the length of the slice, the capacity of the slice, and a reference to the underlying array.

6. With 'value types' in Go, do we have to worry about pointers if we want to pass a value to a function and modify the original value inside the function?
- Yes.

7. With 'reference types' in Go, do we have to worry about pointers if we want to pass a value to a function and modify the original value inside the function?
- No.

8. Is a slice a 'value type' or a 'reference type'
- Reference type, because a slice contains a reference to the actual underlying list of records.

