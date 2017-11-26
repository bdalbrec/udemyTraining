package main

import "fmt"

func main() {
	fmt.Printf("Dec:\tBin:\tHex:\tUtf8: \n")
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
