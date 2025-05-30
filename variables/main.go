package main

import (
	"fmt"
	"reflect"
)

func main() {
	// variable declarations
	var name string = "hello world"
	var isName bool = true
	var number = 123
	floatNumber := 123.132435654321

	fmt.Print("name ", name)
	fmt.Println(" type ", reflect.TypeOf(name))
	fmt.Print("is name ", isName)
	fmt.Println(" type ", reflect.TypeOf(isName))
	fmt.Print("number ", number)
	fmt.Println(" type ", reflect.TypeOf(number))
	fmt.Print("float number ", floatNumber)
	fmt.Println(" type ", reflect.TypeOf(floatNumber))

	// type print techniques
	fmt.Printf("type %T\n", name)
	fmt.Println("type", reflect.TypeOf(name))

	// for switch variable must be of type interface then only type is possible
	var switchName interface{} = name
	switch switchName.(type) {
	case string:
		println("string")
	case bool:
		println("bool")
	default:
		print("others")
	}
}
