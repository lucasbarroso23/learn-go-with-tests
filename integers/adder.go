package main

import "fmt"

// Add takes two integers and returns the sum of then
func Add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(Add(2, 4))
}
