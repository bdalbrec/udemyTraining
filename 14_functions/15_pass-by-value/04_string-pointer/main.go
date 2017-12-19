package main

import "fmt"

func main() {
	name := "Brian"
	fmt.Println(&name)

	changeMe(&name)

	fmt.Println(&name)
	fmt.Println(name) // Rocky
}

func changeMe(z *string) {
	fmt.Println(z)
	fmt.Println(*z) //Brian
	*z = "Rocky"
	fmt.Println(z)
	fmt.Println(*z) //Rocky
}
