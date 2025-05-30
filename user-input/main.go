package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input a number")

	input, err := reader.ReadString('\n')

	fmt.Print(input, " + ", err)
}
