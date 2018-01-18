package main

import "fmt"

func main() {
	// Set up the pipeline
	c := gen(2, 3, 4, 5)
	out := sq(c)

	// Consume the output
	for n := range out { // i think this is because out is a channel so <-out would be a channel of a channel
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
