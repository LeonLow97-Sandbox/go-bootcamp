package main

import (
	"encoding/json"
	"fmt"
)

// Data structure needed to unmarshal (Go in Json)
type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	// JSON
	s := `[{"First":"Leon","Last":"Low","Age":25},{"First":"James","Last":"Bond","Age":32}]`
	// Converting string to byte slice (needed for Unmarshal)
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	// people := []person{}
	var people []person

	err := json.Unmarshal(bs, &people)  // pointed to value
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all of the data", people)
	for i, v := range people {
		fmt.Println(i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}