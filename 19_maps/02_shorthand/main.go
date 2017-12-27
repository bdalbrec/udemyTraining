package main

import "fmt"

func main() {
	myGreeting := map[string]string{ // composite literal
		"Tim":   "Good morning.",
		"Jenny": "Bonjour."}
	fmt.Println(myGreeting)
}
