package main

import "fmt"

var hello string = "hello" //stored in the heap->for persistent(global) memory
func a() {
	var x int = 0 //stored in the stack->for function blocks
	fmt.Printf("%d\n", x)
}
func deallocatingMemory() {

}
