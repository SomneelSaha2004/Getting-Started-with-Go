package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

const maxlength = 20

func main() {
	path := ""
	names := []Name{}
	fmt.Print("Enter file path : ")
	fmt.Scanln(&path)
	f, err := os.Open(strings.TrimPrefix(path, " "))
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
			break
		}
		splitName := strings.Split(string(line), " ")
		newName := Name{fname: splitName[0], lname: splitName[1]}
		names = append(names, newName)
	}
	for i, val := range names {
		fmt.Printf("\n%d\n%s\n%s", (i + 1), SetName(val.fname), SetName(val.lname))
	}
	f.Close()
}
func SetName(s string) string {
	if len(s) > 20 {
		return s[0:21]
	}
	return s
}
