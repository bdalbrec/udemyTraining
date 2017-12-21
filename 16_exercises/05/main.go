package main

import "fmt"

func foo(n ...int) int {
	var biggest int
	for _, x := range n {
		if x > biggest {
			biggest = x
		}
	}
	return biggest
}

func main() {
	n := foo(1, 2, 3)
	fmt.Println(n)
}
