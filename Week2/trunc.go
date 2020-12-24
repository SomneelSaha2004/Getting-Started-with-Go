package main

import "fmt"

func trunc() {
	var f float64
	fmt.Println("Enter float")
	n, err := fmt.Scan(&f)
	if err == nil {
		fmt.Printf("Truncated number = %d\n", int(f))
		fmt.Printf("no of items scanned = %d\n", n)
	} else {
		fmt.Println(err)
	}
}
