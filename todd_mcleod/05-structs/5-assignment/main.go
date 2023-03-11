package main

import "fmt"

type person struct {
	first  string
	last   string
	fav_flavour []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	// exercise1()
	// exercise2()
	// exercise3()
	exercise4()
}

func exercise4() {
	// Create and use an anonymous struct
	p1 := struct {
		first string
		friends map[string]int
		favDrinks []string
	}{
		first: "Leon",
		friends: map[string]int {
			"Leon": 25,
			"Gerrit": 19,
			"Kai Ming": 27,
		},
		favDrinks: []string{
			"Java Chip",
			"Coca Cola",
			"Water",
		},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.friends)
	fmt.Println(p1.favDrinks)

	for k, v := range p1.friends {
		fmt.Println(k, v)
	}

	for i, v := range p1.favDrinks {
		fmt.Println(i, v)
	}
}

func exercise3() {
	// Create a new struct "vehicle"
	// Using composite literal, create a value of type truck and sedan and
	// assign values to the fields
	t1 := truck {
		vehicle: vehicle {
			doors: 2,
			color: "black",
		},
		fourWheel: true,
	}
	t2 := sedan {
		vehicle: vehicle {
			doors: 4,
			color:"orange",
		},
		luxury: false,
	}
	fmt.Println(t1);
	fmt.Println(t2);

	fmt.Println(t1.doors, t1.color, t1.fourWheel)
	fmt.Println(t2.doors, t2.color, t2.luxury)
}

func exercise2() {
	p1 := person{
		first:  "Jie Wei",
		last:   "Low",
		fav_flavour: []string{"chocolate", "martini", "vanilla"},
	}
	p2 := person{
		first:  "Fang Jen",
		last:   "Lim",
		fav_flavour: []string{"hazelnut", "strawnerry", "vanilla"},
	}

	// Store the values of type person in a map with the key of last name
	// Access each value in the map
	// Print out the values, ranging over the slice.
	pMap := map[string]person {
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(pMap)

	for _, value := range pMap {
		fmt.Println(value.first)
		for i, val := range value.fav_flavour {
			fmt.Println(i, val)
		}
	}
}

func exercise1() {
	p1 := person{
		first:  "Jie Wei",
		last:   "Low",
		fav_flavour: []string{"chocolate", "martini", "vanilla"},
	}
	p2 := person{
		first:  "Fang Jen",
		last:   "Lim",
		fav_flavour: []string{"hazelnut", "strawnerry", "vanilla"},
	}

	// Range over the flavours for the 2 structs created
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.fav_flavour {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.fav_flavour {
		fmt.Println(i, v)
	}
}