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
	x := z / 2
	var y bool
	if z%2 == 0 {
		y = true
	} else {
		y = false
	}

	return x, y
}
