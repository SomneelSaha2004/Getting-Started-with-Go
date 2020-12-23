package main

import "fmt"

var y int = 3
var x *int = new(int) //global scope

func f() {
	*x = 3
	fmt.Printf("%d\n", *x)
}
func g() {
	fmt.Printf("%d\n", *x)
}
func scopeOfAVariable() {
	f()
	g()
}

/*
Error code
func f(){
	var x=3->local scope not acessible outside the function
	printf("%d",x)->works perfectly fine
}
func g(){
	printf("%d",x)->does not work as x outside scope
}
*/
