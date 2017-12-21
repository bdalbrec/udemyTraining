package main

import "fmt"

func main() {
	name := "Brian"
	fmt.Println(name) //Brian

	changeMe(name)

	fmt.Println(name) //Brian
}

func changeMe(z string) {
	fmt.Println(z) // Brian
	z = "Rocky"
	fmt.Println(z) // Rocky
}
