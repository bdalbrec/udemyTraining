package main

import "fmt"

func main() {
	b := true
	if food := "Chocolate"; b { // b is what is evaluated, but i can initialize a variable in the statement
		fmt.Println(food)
	}
}
