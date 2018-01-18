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
	fmt.Println(<-c) // we don't range over the channel. We receive the first value then end.
}

//Why does this only print zero?
//And what can you do to get it to printl all 0 - 9 numbers?
