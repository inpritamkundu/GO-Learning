package main

import "fmt"

type User struct {
	name  string
	email string
	id    int
}

func main() {
	userData := User{"pritam", "aadf", 123}

	fmt.Println(userData)
}
