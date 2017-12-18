package main

import "fmt"

func main() {
	greeting := func() {
		fmt.Println("Hello, world!") // this is an anonymous function
	}

	greeting()
}
