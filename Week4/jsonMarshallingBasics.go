package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func jsonMarshallingBasics() {
	//http.Get("www.uci.edu")
	//net.Dial("tcp", "uci.edu:80")
	type Person struct {
		name    string
		address string
		age     int
		phoneNo int64
	}
	p1 := Person{name: "somneel saha", address: "wherever", age: 15, phoneNo: 9811906630}
	jsonPerson, err := json.Marshal(p1)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(jsonPerson)
	var p2 Person
	err = json.Unmarshal(jsonPerson, &p2)
	fmt.Println(p1)
}
