package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // [ ]
	changeMe(m)    // a slice is a REFERECNE type, so no pointer is needed
	fmt.Println(m) // [Brian]
}

func changeMe(z []string) {
	z[0] = "Brian"
	fmt.Println(z) // Brian
}
