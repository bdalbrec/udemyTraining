package main

import "fmt"

func main() {
	var input int
	fmt.Print("Please enter an integer: ")
	fmt.Scan(&input)

	div2, oddOrEven := half(input)
	fmt.Printf("half(%d) returns (%d, %t)", input, div2, oddOrEven)
}

func half(z int) (int, bool) {
	return z / 2, z%2 == 0
}
