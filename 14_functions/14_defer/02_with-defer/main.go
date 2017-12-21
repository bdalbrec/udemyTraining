package main

import "fmt"

func hello() {
	fmt.Print("Hello, ")
}

func world() {
	fmt.Print("world!")
}

func main() {
	defer world()
	hello()
}
