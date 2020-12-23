package main

import "fmt"

func pointerStuff() {
	var x int = 1
	var y int
	var address *int
	address = &x //& operator to find the memory address of a variable
	y = *address //* operator to find the value at a memory address(opposite to &)
	var ptr *float64 = new(float64)
	*ptr = 3
	fmt.Println(*ptr)
	fmt.Println(y)
}
