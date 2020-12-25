package main

import "fmt"

func variableSlices() {
	x := make([]int, 3, 10)
	fmt.Println(x)
	x = append(x, 1)
	fmt.Println(x)
}
