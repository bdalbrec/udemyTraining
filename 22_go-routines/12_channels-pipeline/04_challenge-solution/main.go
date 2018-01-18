package main

import (
	"fmt"
)

func main() {
	// set up the pipeline and consume the output
	for n := 1; n <= 10; n++ {
		for x := 3; x < 13; x++ { // factorial of 100 is way too big for integers so I'm doing a 10 number range 10 times
			fmt.Println("Factorial of ", x, "is:", <-factorial(gen(x)))
		}
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

func factorial(n chan int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := <-n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
