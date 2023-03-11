package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "Leon",
		Last:    "Low",
		Sayings: []string{"Shaken, not stirred", "Any alst wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln("JSON did not marshal:", err)
	}
	fmt.Println(string(bs))
}
