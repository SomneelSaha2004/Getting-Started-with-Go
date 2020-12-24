package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Print("Enter string : ")
	_, err := fmt.Scan(&s)
	if err == nil {
		s = strings.ToLower(s)
		s = strings.TrimPrefix(s, " ")
		if strings.Index(s, "a") != -1 && strings.Index(s, "i") == 0 && strings.LastIndex(s, "n") == len(s)-1 {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	} else {
		fmt.Println("An error has occurred")
		fmt.Println(err)
	}
}
