package main

import (
	"fmt"
)

func main() {
	c := make(chan int) // unbuffered channel blocks until the channel is closed
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done // this is like using an _ for a varialbe to throw it away
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
