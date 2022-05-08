package main

import "fmt"

func main() {
	ar := []int{1, 0, 7, 5, 1, 0, 7}
	finalArr := []int{}
	count := 0
	for _, value := range ar {
		if value == 0 {
			count++
		} else {
			finalArr = append(finalArr, value)
		}
	}
	for i := count; i > 0; {
		finalArr = append(finalArr, 0)
		i--
	}
	fmt.Println("Original Array : ", ar)
	fmt.Println("Final Array : ", finalArr)
}
