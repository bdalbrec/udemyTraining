package main

import "fmt"

func main() {
	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#x \n", 42, 42, 42) //# sign leads leading 0x for hex
	fmt.Printf("%d - %b - %#X \n", 42, 42, 42) // capital X makes upper case
	fmt.Printf("%d \t %b \t %X \n", 42, 42, 42)
}
