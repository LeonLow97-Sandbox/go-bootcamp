package main

// Code that describes what a deck is and how it works

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck' which is a slice of strings
// type 'deck' extends/borrows the behaviors of a slice of string.
type deck []string

// Create and return a list of playing cards. (don't need receiver because we are getting information)
// Returns a value of type deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Receiver Function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal Function (receiver function because we are calling an exiting array)
// (deck, deck) tells Go we want to return 2 values of type 'deck'
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// toString
func (d deck) toString() string {
	// join every value in the slice with a comma in between
	return strings.Join([]string(d), ",")

	// slice of string
	// []string(d)
}

// saveToFile with no return value but an error message
// 0666 : anyone can read and write this file.
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// returns byte slice and error
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Error Handling
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		// quit the program entirely
		os.Exit(1)
	}

	// convert byte slice to string
	s := strings.Split(string(bs), ",")
	// convert string to slice of strings
	return deck(s)
}

func (d deck) shuffle() {
	// generate new seed value
	source := rand.NewSource(time.Now().UnixNano())
	// random number
	r := rand.New(source) // type Rand

	for i := range d {
		// generate a random number
		newPosition := r.Intn(len(d) - 1)

		// swapping the elements at these positions inside the slice
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
