package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	c <- 1 // this is blocking until something takes it off the channel. Need run in a parallel routine so it can be received.
	fmt.Println(<-c)
}

// This results in a deadlock
// Can you determine why?
// And what would you do to fix it?
