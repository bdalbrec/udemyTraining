package main

import "fmt"

func zero(z *int) {
	fmt.Println(z)
	*z = 0 // dereferencing - getting/setting value at memory address
}

func main() {
	x := 5
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x) // x is zero
	fmt.Println(&x)
}
