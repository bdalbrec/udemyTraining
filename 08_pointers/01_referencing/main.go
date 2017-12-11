package main

import (
	"fmt"
)

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a // b is a pointer to an int
	fmt.Println(b)

	// the above code makes b a pointer to the memory address where an int is storec
	//b is of type "int pointer"
	// *int -- the * is part of the type -- b is of type *int
}
