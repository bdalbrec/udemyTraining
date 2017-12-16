package main

import "fmt"

func main() {
	fmt.Println(greet("Jane", "Doe"))
}

func greet(fname, lname string) (string, string) {
	return fmt.Sprint("Hi,", fname, lname), fmt.Sprint("Hi,", lname, fname)
}
