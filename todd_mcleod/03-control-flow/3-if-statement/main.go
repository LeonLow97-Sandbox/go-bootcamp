package main

import "fmt"

func main() {
	ifStatement()
	ifElse()
}

func ifElse() {
	x := 45
	if x == 40 {
		fmt.Println("our value was 40.")
	} else if x == 41 {
		fmt.Println("our value was 41.")
	} else {
		fmt.Println("our value was not 40 or 41.")
	}
}	

func ifStatement() {
	if true {fmt.Println("001")}
	if false {fmt.Println("002")}
	if !true {fmt.Println("003")}
	if !false {fmt.Println("004")}

	if (2 == 2) {fmt.Println("005")}
	if (2 != 2) {fmt.Println("006")}
	if !(2 == 2) {fmt.Println("007")}
	if !(2 != 2) {fmt.Println("008")}

	// Initialization
	if x := 42; x == 42 {
		fmt.Println("009")
	}
	
	// write 2 statements on the same line with semicolon
	fmt.Println("Statement 1"); fmt.Println("Statement 2")
}