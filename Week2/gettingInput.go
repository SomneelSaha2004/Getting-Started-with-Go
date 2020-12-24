package main

import "fmt"

func gettingInput() {
	var x int
	fmt.Println("Enter input")
	num, err := fmt.Scan(&x)
	fmt.Printf("%d : number of items scanned\n", num)
	fmt.Println(err)
	fmt.Println(x)
	fmt.Printf("%T", err)
}
