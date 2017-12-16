package main

import "fmt"

func main() {
	greet("Jane")
	greet("John")
}

func greet(name string) {
	fmt.Println("Hello", name)
}

// greet is declared with a parameter
// when calling greet, pass an argument
