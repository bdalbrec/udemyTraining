package main

import "fmt"

func main() {
	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	// fmt.Println(food) // can't do it. Out of scope.
}
