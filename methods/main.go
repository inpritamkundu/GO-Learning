package main

import "fmt"

func main() {
	// when a function is passed in struct it is called methods
	firstUser := User{"aman", "aman@mailinator.com"}
	firstUser.emailPrint()
}

type User struct {
	name  string
	email string
}

func (u User) emailPrint() {
	fmt.Println(u.email)
}
