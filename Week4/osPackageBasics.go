package main

import (
	"fmt"
	"os"
)

func osPackageBasics() {
	arr := make([]byte, 100)
	file, err := os.Open("more.txt")
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Read(arr)
	fmt.Println(string(arr))
	file.Close()
	file, err = os.Create("out.txt")
	arr = []byte{1, 2, 3}
	_, err = file.Write(arr)
	_, err = file.WriteString("Hello World")
	file.Close()
}
