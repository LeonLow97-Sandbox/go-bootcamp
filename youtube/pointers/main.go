package main

import "fmt"

type Player struct {
	health int
}

type BigData struct {
	// 500 MB
}

// Real life scenario: Calling 100000 times (too troublesome so use pointers)
// Use 8 bytes => pointer (64 bit machines), better performance
func processBigData(bd BigData) {
	
}

func (player *Player) takeDamageFromExplosion(dmg int) { // same function
	fmt.Println("Player is taking damage from explosion!")
	player.health -= dmg
}

func takeDamageFromExplosion(player *Player, dmg int) { // same function
	fmt.Println("Player is taking damage from explosion!")
	player.health -= dmg
	// return
}

func main() {
	// 8 byte long integer (pointer)
	player := &Player{
		health: 100,
	}

	fmt.Printf("Before explosion %+v\n", player)
	takeDamageFromExplosion(player, 10) // copy
	player.takeDamageFromExplosion(10)
	fmt.Printf("After explosion %+v\n", player)
}
