package main

// Code to create and manipulate a deck
func main() {
	// declare a slice
	// deck --> []string
	// cards is of type "deck" so can call "print()" method
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// cards.print()

	// cards := newDeck()

	// // because we are returning 2 values from deal() function
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	// cards.shuffle()
	cards.print()
}

// specify the data type that we are going to return from any function
// func newCard() string {
// 	return "Five of Diamonds"
// }
