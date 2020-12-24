package main

import "fmt"

func typesAndConversion() {
	var x int16 = 1
	var y int32 = 2
	//x=y is an error as they are of different types
	x = int16(y)
	fmt.Printf("%d\n", x)
	var s string = fmt.Sprint(x)
	fmt.Println(s)
	var f float64 = 1.2345e2 //e can be used to represent 10^2
	fmt.Printf("%f", f)
	var c complex128 = complex(2, 3)
	fmt.Println(c)
}
