package main

import "fmt"

func main() {
	if true && true {
		fmt.Println("This printed")
	}

	if true && false {
		fmt.Println("This will not print")
	}

	if false && false {
		fmt.Println("This will not print")
	}

}
