package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(sf ...float64) float64 { // the ...float64 means it takes 0 or more float64 arguments
	var total float64
	fmt.Println(sf)
	fmt.Printf("%T\n", sf)
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
