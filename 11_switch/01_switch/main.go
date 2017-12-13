package main

import "fmt"

func main() {
	switch "John" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have you no friends?")
	}
}

/*
  no default fallthrough
  -- you can specify fallthrough by explicitly stating it
  -- break isn't needed like in other languages
*/
