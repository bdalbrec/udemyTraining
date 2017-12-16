package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57} // this is one item, a slice
	n := average(data...)                     // opens the single item and passes in the values to the variatic function
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
