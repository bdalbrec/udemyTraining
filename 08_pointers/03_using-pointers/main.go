package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x some stuff

	var b *int = &a
	fmt.Println(b)  // 0x some stuff
	fmt.Println(*b) //43

	*b = 42        // assigns 42 to the address b points to
	fmt.Println(a) //42

	// this is useful
	// we can pass a memory address instead of a bunch of values (we can pass a reference)
	// and then we can still change the value of whatever is stored at that memory address
	// this makes our programs more performant
	// we don't have to pass around large amounts of data
	// we only have to pass addresses
}
