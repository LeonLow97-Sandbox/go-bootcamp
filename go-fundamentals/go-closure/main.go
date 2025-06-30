package main

import "fmt"

func activateGiftCard() func(int) int {
	// Closure
	// Go will store any variables that are defined outside (`amount`)
	// of the closure function in memory
	amount := 100 // not erased from memory because we have closure (`debitFunc`)
	debitFunc := func(debitAmount int) int {
		amount -= debitAmount
		return amount
	}

	return debitFunc
}

func main() {
	// activate 2 gift cards
	useGiftCard1 := activateGiftCard()
	useGiftCard2 := activateGiftCard()

	fmt.Println(useGiftCard1(10))
	fmt.Println(useGiftCard2(20))
}
