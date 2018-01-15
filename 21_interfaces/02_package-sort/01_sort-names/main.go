package main

import (
	"fmt"
	"sort"
)

type people []string // implements Interface interface

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(studyGroup)
	sort.Sort(studyGroup) // sort takes a type of Interface as an argument
	fmt.Println(studyGroup)
}
