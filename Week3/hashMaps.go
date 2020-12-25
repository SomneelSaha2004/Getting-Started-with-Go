package main

import "fmt"

func hashMaps() {
	name := map[string]string{"first": "somneel", "middle": "kumar"}
	fmt.Println(name["first"])
	name["last"] = "saha"
	fmt.Println(name)
	delete(name, "first")
	fmt.Println(len(name))
	id, n := name["middle"]
	fmt.Println(id, n)
	for id, val := range name {
		fmt.Printf("Key : %s	Value : %s\n", id, val)
	}
}
