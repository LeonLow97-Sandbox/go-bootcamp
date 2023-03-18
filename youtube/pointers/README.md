# Pointers in Golang

## Why should we use pointers? And when?

- Using pointers in Go can **reduce memory usage**  by avoiding the need to make copies of large or complex data structures.
- When a value is passed as an argument to a function in Go, a copy of that value is created and passed to the function.
  - This can be inefficient for large or complex data structures, as it can consume a significant amount of memory and slow down your program.
- Pointers allow you to pass a reference to a value instead of copy of the value itself.
  - This can help to reduce memory usage, as you are only passing a small pointer instead of the entire value.
- When you use a pointer to access a value, you are modifying the value in place instead of making a copy of the value and returning a new modified value.
  - This can be more memory-efficient, as you are not creating multiple copies of the same value.
- Pointers can be useful for sharing data between functions or goroutines, as you can pass a pointer to a shared value and allow multiple functions or goroutines to access and modify the value in place.

## Overall

- Using pointers in Go can be an efficient way to reduce memory usage and improve performance for large or complex values.
- However, it's important to use pointers carefully and avoid common pitfalls such as dereferencing a nil pointer or leaking memory by not properly freeing allocated memory.

## Examples of large data structures in Golang

1. Arrays and Slices
2. Maps
3. Structs
4. Trees and Graphs