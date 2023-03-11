package main

import "fmt"

func main() {
	// Maps are an unordered list (key-value pair)
	// Creating a map
	m := map[string]int{
		"Leon":           25,
		"Ms Fang Jen": 23,
	}
	fmt.Println(m)
	fmt.Println(m["Leon"])
	fmt.Println(m["Jie En"]) // entering a key that does not exist returns value 0

	v, ok := m["Jie En"]
	fmt.Println(v); fmt.Println(ok)

	if v, ok := m["Jie En"]; ok {
		fmt.Println("This is the IF print", v)
	}

	// Add element in Map
	m["Tuck Yern"] = 25
	fmt.Println(m)

	for keys, values := range m {
		fmt.Println(keys, values)
	}

	// Delete element in Map
	// delete(<map name>, "key")
	delete(m, "Tuck Yern")
	fmt.Println(m)

	// Delete element that does not exist
	delete(m, "Nicholas")
	fmt.Println(m) // no error returned

	// To ensure that something is deleted
	if v, ok := m["Leon"]; ok {
		fmt.Println("value", v)
		delete(m, "Leon")
	}
	fmt.Println(m)
}