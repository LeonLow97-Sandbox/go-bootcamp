package main

import "fmt"

func main() {
	users()

	grades()
}

func grades() {
	gradesMap := make(map[string]int)

	gradesMap = map[string]int {
		"Leon": 89,
		"Darrel": 100,
	}

	for keys := range gradesMap {
		fmt.Println(keys)
	}
}

func users() {
	userMap := make(map[string]string)

	userMap["Leon"] = "Low Jie Wei"
	userMap["Darrel"] = "Ang Jing Sheng"
	userMap["Bella"] = "Low Jie Xin"

	// checking for existence
	user, exists := userMap["Leon"]
	if exists {
		fmt.Println(user)
	} else {
		fmt.Println("user does not exist")
	}

	// iterating over map
	for key, value := range userMap {
		fmt.Println(key, ":", value)
	}

	// deleting a key-value pair
	delete(userMap, "Leon")
	if value, exists := userMap["Leon"]; !exists {
		fmt.Println("user does not exist")
	} else {
		fmt.Println(value, "exists")
	}
}
