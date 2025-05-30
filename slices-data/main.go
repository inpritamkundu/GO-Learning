package main

import (
	"fmt"
	"sort"
)

func main() {

	number := []int{}
	fmt.Println("number init", number)

	// add data || append data

	number = append(number, 5, 6)
	fmt.Println("number add", number)

	// slice data

	number = number[1:]
	fmt.Println("number removed", number)

	// make slice using make

	numbersMake := make([]int, 3)
	numbersMake[0] = 5
	numbersMake[1] = 2
	numbersMake[2] = 3
	fmt.Println("array using make", numbersMake)

	// sort slice
	sort.Ints(numbersMake)
	fmt.Println("array sort", numbersMake)

	// delete data through index
	index := 1
	numbersMake = append(numbersMake[:index], numbersMake[index+1:]...)
	fmt.Println("index deleted", numbersMake)

}
