package main

import "fmt"

func printNames(names ...string) {
	for index, name := range names {
		fmt.Println(index, name)
	}
}

func main() {
	students := []string{"Ashok", "Alvin", "Simon"}
	printNames(students...)
}
