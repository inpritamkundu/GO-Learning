package main

import "fmt"

func main() {
	m := make(map[int]string)

	m[0] = "hello"
	m[1] = "world"
	m[2] = "you"
	m[3] = "are"

	fmt.Println(m)

	// delete element on basis of key

	delete(m, 0)
	fmt.Println(m)

	// length of map
	length := len(m)

	fmt.Println("length", length)

	// looping map

	for key, val := range m {
		fmt.Println(key, val)
	}

}
