package main

import "fmt"

func add(a, b int) int {
	c := a + b
	return c
}

func main() {
	result := add(2, 9)
	fmt.Println(result)
}
