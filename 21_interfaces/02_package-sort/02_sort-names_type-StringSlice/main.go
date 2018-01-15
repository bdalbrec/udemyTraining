package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	// sort.StringSlice(s).Sort()  // converts to StringSlice type which has a sort method
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
}
