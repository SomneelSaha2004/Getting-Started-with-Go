package main

import "fmt"

//const keyword used to define constant values in the code
func constants() {
	const x int = 3
	const (
		y int     = 4
		z float64 = 30.222
	)
	const sum float64 = float64(x+y) + z
	fmt.Printf("%f", sum)
	const b = 1
	type Grades int
	const (
		A Grades = iota
		B
		C
		D
		E
		F
	)
}
