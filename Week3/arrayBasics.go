package main

import "fmt"

func arrayBasics() {
	var x [5]int
	x[0] = 2
	fmt.Println(x[1])
	fmt.Println(x)
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println(len(arr))
	for i, val := range arr {
		fmt.Printf("Index : %d	Value: %d\n", i, val)
	}
}
