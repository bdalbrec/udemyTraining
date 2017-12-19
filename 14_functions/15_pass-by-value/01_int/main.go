package main

import "fmt"

func main() {
	age := 44
	fmt.Println(&age)

	changeMe(&age)

	fmt.Println(&age)
	fmt.Println(age) //24
}

func changeMe(z *int) {
	fmt.Println(z)
	fmt.Println(*z) //44
	*z = 24

	fmt.Println(z)
	fmt.Println(*z)
}
