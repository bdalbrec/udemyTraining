package main

import "fmt"

func findGreatest(si ...int) int {
	var biggest int
	for _, i := range si {
		if i > biggest {
			biggest = i
		}
	}
	return biggest
}

func main() {
	fmt.Println(findGreatest(1, 2, 3, 4, 5, 18, 6, 11))
}
