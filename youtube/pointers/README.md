# Pointers

- In this example, we will be working with the `Player` struct.

```go
type Player struct {
	health int
}
```

## Copy

- In the example below, the `Player` struct passed to the `takeDamage` function is a **copy** of the original.
- Any modifications to `player` in `takeDamage` will not affect the original `Player` struct in the `main` function.
- Hence, the player's health after taking damage will remain as 100.

```go
func takeDamage(player Player) {
	dmg := 10
	player.health -= dmg
}

func main() {
	// 1. Copy
	player := Player{
		health: 100,
	}
	takeDamage(player)
	fmt.Println("Players Health (Copy): ", player.health) // Output: Players Health (Copy):  100
}
```

## Functional

- The functional method receives a copy of the `Player` struct as its receiver, so any modifications made to it won't affect the original `Player` instance.
- It subtracts the damage from the copy and returns a new `Player` instance with the updated health.
- The method creates a returns a new `Player` struct, which **consumes additional memory**, but it doesn't increase memory usage for the original `Player` instance since it operates on a copy.
- The new `Player` struct is then reassigned to the `player` variable in the `main` function.

```go
func (player Player) takeDamageFunctional() Player {
	dmg := 10
	player.health -= dmg
	return player
}

func main() {
    player = Player{
		health: 100,
	}
	player = player.takeDamageFunctional()
	fmt.Println("Players Health (Functional): ", player.health) // Output: Players Health (Functional):  90
}
```

## Pointer

- The method with a pointer-receiver works on the original `Player` instance because it receives a pointer to the `Player` struct. Any changes made to the receiver inside the method will directly modify the original `Player` instance.
- It subtracts the damage from the `Player` instance, effectively modifying its health directly without returning anything.
- Since it modifies the original `Player` instance directly, it doesn't need to create a new instance or allocate additional memory. It updates the existing `Player` instance in place.

```go
// Method 1 - Pointer and Receiver
func (player *Player) takeDamagePointer() {
	dmg := 10
	player.health -= dmg
}

player = Player{
    health: 100,
}
player.takeDamagePointer() // not &player.takeDamagePointer().
fmt.Println("Players Health (Method 1 - Pointer and Receiver): ", player.health)

// Method 2 - Pointer
func takeDamagePointer2(player *Player) {
	dmg := 10
	player.health -= dmg
}

func main() {
	player = Player{
		health: 100,
	}
	takeDamagePointer2(&player) // passing the memory address of the Player struct with `&`
	fmt.Println("Players Health (Method 2 - Pointer and Receiver): ", player.health)
}
```

# Comparison

---

### Functional Approach

The functional approach emphasizes immutability, and thus, operations typically return a new value rather than modifying an existing one.

#### Pros

1. _Safety_: It avoids the risk of unintended side effects because data is not mutated.
2. _Testability_: It is often easier to test functional code due to its deterministic nature.
3. _Concurrency_: In multi-threaded environments, usually don't need to worry about race conditions because you are not mutating shared state.

#### Cons

1. _Performance_: Creating a new copy of large structures every time can lead to performance overhead.
2. _Memory Usage_: The need to frequently create copies can lead to higher memory usage.

---

### Pointer Approach

The pointer approach allows you to directly modify the original data structure.

---

#### Pros

1. _Efficiency_: Pointers provide a more efficient way to handle large data structures as you are not creating a copy.
2. _Control_: You have more control over the memory you are working with.

#### Cons

1. _Mutability_: There is a risk of mutating data unintentionally, which can cause bugs.
2. _Concurrency Issues_: In multi-threaded environments, it multiple threads are mutating a shared state simultaneously, it can lead to race conditions.

```go
// causes issue if somewhere in the code, garbage collection is performed on the variable but
// subsequent code is still relying on the Player struct
player = nil
printPlayerHealth(&player.health) // Output: panic: runtime error: invalid memory address or nil pointer dereference

// What a nasty error message!
```

---

### Functional Approach or Pointer Approach?

1. If you are dealing with small data structures or prioritizing the benefits of immutability such as safety, testability, and simplicity, especially in a multi-threaded environment, then a `functional` approach might be more suitable.
2. If you are working with large data structures where copying would be too expensive, or if you require the efficiency and control that pointers provide, then the `pointer` approach might be better.
