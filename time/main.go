package main

import (
	"fmt"
	"time"
)

func main() {

	presentTime := time.Now()

	fmt.Println("date format: ", presentTime)

	// the numbers for the format are fixed

	// 	Month: 1 (January)
	// Day: 2
	// Hour: 3 (15 in 24-hour format)
	// Minute: 4
	// Second: 5
	// Year: 6 (2006)
	// Timezone: -0700

	fmt.Println("formatting date: ", presentTime.Format("02-01 15:04"))
}
