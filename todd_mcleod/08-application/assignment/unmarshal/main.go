package main

import (
	"encoding/json"
	"fmt"
)

type Person []struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Age":32},{"First":"Moneypenny","Age":27},{"First":"M","Age":54}]`
	bs := []byte(s)

	var people Person

	// func Unmarshal(data []byte, v any) error
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)

	for _, v := range people {
		fmt.Println("Person", v.First, v.Age)
	}
}