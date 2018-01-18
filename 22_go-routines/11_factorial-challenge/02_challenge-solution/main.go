package main

import (
	"fmt"
)

func main() {
	f := factorial(4)
	fmt.Println("Total:", <-f)
}

func factorial(n int) chan int {
	out := make(chan int)
	total := 1
	go func() {
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
