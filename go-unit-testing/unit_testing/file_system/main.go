package main

import (
	"fmt"
	"net/http"

	"github.com/LeonLow97/handlers"
)

func main() {
	http.HandleFunc("/view", handlers.ReadFile)

	fmt.Println("Server is listening on :8080!")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
