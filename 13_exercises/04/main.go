package main

import "fmt"

func main() {

	var smallNumber int
	var largeNumber int
	fmt.Print("Please enter a small number: ")
	fmt.Scan(&smallNumber)
	fmt.Print("Please enter a large number: ")
	fmt.Scan(&largeNumber)
	fmt.Printf("%d divided by %d has a remainder of %d.\n", largeNumber, smallNumber, largeNumber%smallNumber)
}
