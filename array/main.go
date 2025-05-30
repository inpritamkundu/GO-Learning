package main

import "fmt"

func main() {
	var stringArray [3]string

	stringArray[0] = "a"
	stringArray[2] = "d"

	fmt.Println(stringArray)

	numberArray := [3]int{1, 2, 3}

	fmt.Println(numberArray)

	Number2DArray := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println(Number2DArray)
}
