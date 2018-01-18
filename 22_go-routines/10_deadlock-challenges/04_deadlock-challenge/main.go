package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		} // the channel is never closed. That is one problem.
	}()

	for {
		fmt.Println(<-c) // this will also cause an issue. Needs to be a range so the loop ends instead of going infinitely.
	}
}

// This still deadlocks after printing out all the numbers on the channel.
