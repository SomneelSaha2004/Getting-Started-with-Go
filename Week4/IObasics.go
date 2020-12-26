package main

import (
	"fmt"
	"io/ioutil"
)

func IObasics() {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	s := string(dat)
	fmt.Printf("Writing to output file -----> %s", string(dat))
	err = ioutil.WriteFile("output.txt", []byte(s), 0777)
}
