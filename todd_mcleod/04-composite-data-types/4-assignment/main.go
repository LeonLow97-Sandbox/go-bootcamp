package main

import "fmt"

func main() {
	// exercise1()
	// exercise2()
	// exercise3()
	// exercise4()
	// exercise5()
	// exercise6()
	// exercise7()
	// exercise8()
	// exercise9()
	exercise10()
}

func exercise10() {
	x := map[string][]string{
		"bond_james": []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr": []string{"Being evil", "Ice cream", "Sunsets"},
	}
	// Add a record to the map and print out using "range" loop
	x["leon"] = []string{"NUS", "Male", "Funny"}

	// Delete a record from the map
	delete(x, "leon")
	
	for keys, values := range x {
		fmt.Println("This is the record for", keys)
		// for loop over the slice of string
		for index, value := range values {
			fmt.Println("\t", index, value)
		}
	}
}

func exercise9() {
	x := map[string][]string{
		"bond_james": []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr": []string{"Being evil", "Ice cream", "Sunsets"},
	}
	// Add a record to the map and print out using "range" loop
	x["leon"] = []string{"NUS", "Male", "Funny"}

	for keys, values := range x {
		fmt.Println("This is the record for", keys)
		// for loop over the slice of string
		for index, value := range values {
			fmt.Println("\t", index, value)
		}
	}
}

func exercise8() {
		// Create a map with a key of type string and value of type string
		x := map[string][]string{
			"bond_james": []string{"Shaken, not stirred", "Martinis", "Women"},
			"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
			"no_dr": []string{"Being evil", "Ice cream", "Sunsets"},
		}

		for keys, values := range x {
			fmt.Println("This is the record for", keys)
			// for loop over the slice of string
			for index, value := range values {
				fmt.Println("\t", index, value)
			}
		}
}

func exercise7() {
	// Create a slice of a slice of string (multi-dimensional slice)
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Hello, James"}
	md := [][]string{jb, mp}
	// fmt.Println(md)

	for _, values1 := range md {
		for _, values2 := range values1 {
			fmt.Println(values1, values2)
		}
	}
}

func exercise6() {
	// Creating slice using make
	s := make([]string, 50, 50)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	states := []string{"Alabama","Alaska","Arizona","Arkansas","California","Colorado","Connecticut","Delaware","Florida","Georgia","Hawaii","Idaho","Illinois","Indiana","Iowa","Kansas","Kentucky","Louisiana","Maine","Maryland","Massachusetts","Michigan","Minnesota","Mississippi","Missouri","Montana","Nebraska","Nevada","New Hampshire","New Jersey","New Mexico","New York","North Carolina","North Dakota","Ohio","Oklahoma","Oregon","Pennsylvania","Rhode Island","South Carolina","South Dakota","Tennessee","Texas","Utah","Vermont","Virginia","Washington","West Virginia","Wisconsin","Wyoming"}
	fmt.Println(len(states))

	// cannot use append as it will append to the end of the slice make
	// use for loop to store all the states in the "make" slice
	for i, v := range states {
		s[i] = v
	}

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}
}

func exercise5() {
	// To delete from a slice, use APPEND along with SLICING
	x := []int{42,43,44,45,46,47,48,49,50,51}
	// Use append and slicing to get [42, 43, 44, 48, 49, 50, 51]
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}

func exercise4() {
	x := []int{42,43,44,45,46,47,48,49,50,51}

	// Append 52 to slice x
	x = append(x, 52)
	fmt.Println(x)

	// In one statement, append several values
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	// Append to the slice with another slice
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

func exercise3() {
	// Use slicing the create new slices
	x := []int{42,43,44,45,46,47,48,49,50,51}
	// [42 43 44 45 46]
	fmt.Println(x[:5])

	// [47 48 49 50 51]
	fmt.Println(x[5:])

	// [44 45 46 47 48]
	fmt.Println(x[2:7])

	// [43 44 45 46 47]
	fmt.Println(x[1:6])
}

func exercise2() {
	// Create a slice of type int and assign 10 values
	x := []int{42,43,44,45,46,47,48,49,50,51}
	for _, values := range x {
		fmt.Println(values)
	}
	fmt.Printf("%T\n", x)
}

func exercise1() {
	// Create an array that holds 5 values of type INT
	// Print out all values and type of array
	x := [5]int{32,43,54,67,78}

	for index, value := range x {
		fmt.Println(index, value)
	}

	fmt.Printf("%T\n", x)
}