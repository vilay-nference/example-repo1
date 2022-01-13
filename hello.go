package main

import "fmt"

func main() {

	// Get the sum of two number
	var s int = sum(4, 5)
	fmt.Println("Hello world", s)
}

func sum(a, b int) int {
	return a + b
}
