package main

import "fmt"

func main() {
	// defer means the statement or function or expression will execute at last of the main block with LIFO concept

	defer fmt.Println("hello")
	defer fmt.Println("world")
	defer fmt.Println("one")

	fmt.Println("hello boy")
	deferred()

	// so the stack should print like

	// hello world one hello boy 0 1 2 3 4

	// but due to defer
	// output : hello boy 4 3 2 1 0 one world hello
}

func deferred() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
