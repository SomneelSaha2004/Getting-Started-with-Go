package main

import "fmt"

func foo() *int {
	var x int = 3
	return &x
}
func garbageCollection() {
	var y *int
	y = foo()
	fmt.Printf("%d", *y)
}
