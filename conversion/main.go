package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var num string = "3"

	convertedNumber, err := strconv.ParseFloat(num, 32)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(convertedNumber, reflect.TypeOf(convertedNumber))

}
