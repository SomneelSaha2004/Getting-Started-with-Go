package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Enter X to exit program")
	nums := make([]int, 3)
	i := 2
	for true {
		var input string
		fmt.Print("Enter integer : ")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("An error has occurred")
			fmt.Println(err)
			break
		}
		if input == "X" {
			sort.Ints(nums)
			fmt.Println(nums)
			fmt.Println("Program has terminated")
			break
		} else {
			x, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("An error has occurred")
				fmt.Println(err)
				break
			}
			if i == -1 {
				nums = append(nums, x)
				sort.Ints(nums)
				fmt.Println(nums)
				continue
			}
			nums[i] = x
			sort.Ints(nums)
			fmt.Println(nums, i)
			i--
		}
	}
}
