package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (a ByLast) Len() int           { return len(a) }
func (a ByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

// Sort the []user by `age` and `last`
// Sort each []string "Sayings" for each user
// Print everything out using the range clause
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
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

	fmt.Println("\nSorting by Age...")

	// Sort by age
	sort.Sort(ByAge(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

	fmt.Println("\nSorting by Last...")

	// Sort by last
	sort.Sort(ByLast(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

	// Sort by Sayings
	fmt.Println("Sorting by Sayings...")
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		// Sorting by Sayings
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}
}