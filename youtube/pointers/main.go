package main

import "fmt"

type Player struct {
	health int
}

// Makes a copy of player
func takeDamage(player Player) {
	dmg := 10
	player.health -= dmg
}

// Functional - returns a new Player struct
func (player Player) takeDamageFunctional() Player {
	dmg := 10
	player.health -= dmg
	return player
}

// Method 1 - Pointer and Receiver
func (player *Player) takeDamagePointer() {
	dmg := 10
	player.health -= dmg
}

// Method 2 and 3 - Pointer and Receiver
func takeDamagePointer2(player *Player) {
	dmg := 10
	player.health -= dmg
}

func main() {
	// 1. Copy
	player := Player{
		health: 100,
	}
	takeDamage(player)
	fmt.Println("Players Health (Copy): ", player.health)

	// 2. Functional
	player = Player{
		health: 100,
	}
	player = player.takeDamageFunctional()
	fmt.Println("Players Health (Functional): ", player.health)

	// 3. Method 1 - Pointer and Receiver (automatic conversion between a value receiver and a pointer receiver)
	player = Player{
		health: 100,
	}
	player.takeDamagePointer() // not &player.takeDamagePointer().
	fmt.Println("Players Health (Method 1 - Pointer and Receiver): ", player.health)

	// 3. Method 2 - Pointer (passing address into function)
	player = Player{
		health: 100,
	}
	takeDamagePointer2(&player)
	fmt.Println("Players Health (Method 2 - Pointer and Receiver): ", player.health)

	// 3. Method 3 - Pointer (passing address into function by assigning `&` to struct)
	player2 := &Player{
		health: 100,
	}
	takeDamagePointer2(player2)
	fmt.Println("Players Health (Method 3 - Pointer and Receiver): ", player.health)
}
