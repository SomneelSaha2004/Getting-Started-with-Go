package main

import "fmt"

func slicesBasics() {
	arr := [...]string{"a", "b", "c", "d", "e", "f"}
	x2 := arr[1:3] //inclusive of start and exclusive of end
	x1 := arr[2:5]
	fmt.Println(x1, x2)
	x1[0] = "somneel"
	fmt.Println(cap(x1), len(x1))
	fmt.Println(arr, x1, x2)
	slice := []int{1, 2, 3}
	fmt.Println(slice[2])
}
