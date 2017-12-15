package main

import "fmt"

func main() {
	var userName string
	fmt.Print("Please enter your name: ")
	fmt.Scan(&userName)
	fmt.Printf("Hello %s!\n", userName)
}
