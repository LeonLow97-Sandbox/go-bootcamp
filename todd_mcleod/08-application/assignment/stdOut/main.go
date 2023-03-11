package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// Encode to JSON the []user sending the results to Stdout.
// Need to use `json.NewEncoder(os.Stdout).encode(v interface{})`
func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}
	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself",
		},
	}
	u3 := user{
		First: "Eudora",
		Last:  "Leon",
		Age:   25,
		Sayings: []string{
			"Would you like to have a drink with me?",
			"The more your fear grows, the smaller you become",
			"Give life to magic, not magic to life",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// Turning data structure to JSON
	err := json.NewEncoder(os.Stdout).Encode(users)

	if err != nil {
		fmt.Println("We did something wrong and here's the error:", err)
	}

	
}