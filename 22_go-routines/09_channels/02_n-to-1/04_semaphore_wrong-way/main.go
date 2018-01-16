package main

import (
	"fmt"
)

func main() {
	c := make(chan int) // unbuffered channel blocks until the channel is closed
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // has to have a receiver open to proceed. program deadlocks
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // cam't be removed from channel because the range to receive it is blocked
		}
		done <- true
	}()

	// we block here until done <- true
	<-done
	<-done
	close(c)

	// to unblock above
	// we need to take values off of chan c here
	// but we never get here, because we're blocked above
	for n := range c {
		fmt.Println(n)
	}
}
