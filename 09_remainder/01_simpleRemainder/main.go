package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {
		if i%2 == 1 {
			fmt.Println(i, "Odd")
		} else {
			fmt.Println(i, "Even")
		}
	}
}
