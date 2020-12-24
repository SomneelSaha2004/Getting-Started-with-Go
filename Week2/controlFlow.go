package main

import "fmt"

func controlFlow() {
	var x int = 4.0
	if x%2 == 0 {
		fmt.Printf("%d is an even number", x)
	} else {
		fmt.Printf("%d is an odd number", x)
	}
	for i := 1; i <= 10; {
		fmt.Println(i)
		if i == 5 {
			fmt.Println("Peace out!")
			break
		}
		i++
	}
	const i int = 1
	switch i {
	case 1:
		fmt.Println("Hey")
	case 2:
		fmt.Println("hi")
	case 3:
		fmt.Println("Hello")
	default:
		fmt.Println("Wrong value has been passed")
	}
	switch {
	case x > 3:
		fmt.Println("Hello >3")
	case x < 0:
		fmt.Println("Hello negative")
	case x > 10:
		fmt.Println("Hello >10")
	}

}

/*
if <cond>{
	<statements>
}
for <init>;<cond>;<update>{
	<statements>
}
*/
