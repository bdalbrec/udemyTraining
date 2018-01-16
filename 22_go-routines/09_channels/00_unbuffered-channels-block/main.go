package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // this is an unbuffered channel could be make(chan int, 10) to transfer 10 at once

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)
}
