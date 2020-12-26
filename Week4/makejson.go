package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func makejson() {
	reader := bufio.NewReader(os.Stdin)
	x := make(map[string]string)
	fmt.Print("Enter name : ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error has occurred\n", err)
		os.Exit(1)
	}
	//fmt.Println(name)
	x["name"] = name
	fmt.Print("Enter address : ")
	address, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error has occurred\n", err)
		os.Exit(1)
	}
	x["address"] = address
	//fmt.Println(address)
	jsonObj, err := json.Marshal(x)
	if err != nil {
		fmt.Println("An error has occurred\n", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonObj))
}
