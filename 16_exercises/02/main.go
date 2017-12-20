package main

import "fmt"

func main() {
	var input int
	fmt.Print("Please enter an integer: ")
	fmt.Scan(&input)

	calculate := func() {
		x := input / 2
		var y bool
		if input%2 == 0 {
			y = true
		} else {
			y = false
		}
		fmt.Printf("half(%d) returns (%d, %t)", input, x, y)
	}
	calculate()
}
