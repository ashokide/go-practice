package main

import (
	"fmt"
	"practice/students"
)

type Student struct {
	name    string
	marks   []float32
	total   float32
	average float32
}

func printAnalysis(s Student) {
	fmt.Println("Name: " + s.name)
	fmt.Println("Marks: ", s.marks)
	fmt.Println("Total: ", s.total)
	fmt.Println("Average: ", s.average)
}

func main() {
	student1 := Student{
		name:  "Ashok",
		marks: []float32{68, 72, 54, 39, 89},
	}
	student1.total = students.Total(student1.marks)
	student1.average = students.Average(student1.marks)
	printAnalysis(student1)
}
