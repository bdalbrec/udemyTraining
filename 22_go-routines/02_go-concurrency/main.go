package main

import "fmt"

func main() {
	go foo()
	go bar()
	// nothing prints because we aren't waiting for the functions to run before hitting the end of main()
}

func foo() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Bar:", i)
	}
}
