package main

import "fmt"

func main() {
	hello()

	returnVal := sum(2, 5)
	fmt.Println(returnVal)

	valuesSum, textVal := sumOfUnlimitedValues(5, 76, 4, 32, 1, 456, 7)
	fmt.Println(valuesSum, textVal)
}

func hello() {
	fmt.Println("hello world")
}

func sum(a int, b int) int {
	return a + b
}
func sumOfUnlimitedValues(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total = total + val
	}

	return total, "string can also be returned"
}
