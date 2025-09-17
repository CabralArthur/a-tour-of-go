package main

import "fmt"

// When two or more consecutive function parameters share the same type,
// you can omit the type name for the like-typed parameters up to the final parameter that declares the type.
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
