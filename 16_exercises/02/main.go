package main

import "fmt"

func main() {
	var input int
	fmt.Print("Please enter an integer: ")
	fmt.Scan(&input)

	calculate := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}
	fmt.Println(calculate(input))
}
