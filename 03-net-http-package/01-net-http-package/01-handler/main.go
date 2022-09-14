package main

import (
	"fmt"
	"net/http"
)

type jiewei int

// Attach to method serveHTTP
func (j jiewei) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Jie Wei, you can insert your code here!")
}

// handles any path
func main() {
	var d jiewei
	http.ListenAndServe(":8080", d)
}
