package main

import "fmt"

func main() {
	student := []string{}
	students := [][]string{}
	//student[0] = "Brian" // won't work as the length of student is 0
	student = append(student, "Brian")
	fmt.Println(student)
	fmt.Println(students)
}
