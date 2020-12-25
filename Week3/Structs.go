package main

import "fmt"

func Structs() {
	type Person struct {
		name    string
		address string
		number  int64
	}
	var p1 Person
	p1.name = "Somneel"
	p1.address = "work"
	p1.number = 1123231232
	fmt.Println(p1)
}
