package main

import (
	"fmt"
)

type ninjaStar struct {
	owner string
}

type ninjaSword struct {
	owner string
}

func (ninjaStar) attack() {
	fmt.Println("Throwing ninja star")
}

func (ninjaStar) pickup() {
	fmt.Println("Picking up ninja star")
}
 
func (ninjaSword) attack() {
	fmt.Println("Throwing ninja sword")
}

type ninjaWeapon interface {
	attack()
}

func main() {
	weapons := []ninjaWeapon {
		ninjaStar{"Wallace"},
		ninjaSword{"Wallace"}, // panic on weapon.(ninjaStar).pickup()
	}

	// weapon.attack() is of type ninjaWeapon interface but we want it to be 
	// type ninjaStar to access the pickup method. use type assertion!
	for _, weapon := range weapons {
		weapon.attack()

		// Error handling for type assertion
		ns, ok := weapon.(ninjaStar)
		if ok {
			ns.pickup()
		}
	}

	var i interface{} = "Hello world!"
	s, ok := i.(int)
	if ok {
		fmt.Println(s)
	} else {
		fmt.Printf("The provided interface is not of type integer but of type %T\n", i)
	}

}