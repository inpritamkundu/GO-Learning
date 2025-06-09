package main

import "fmt"

func main() {
	weekDays := []string{"sunday", "monday", "tuesday"}

	fmt.Println(weekDays)

	for i := 0; i < len(weekDays); i++ {
		fmt.Println(weekDays[i])
	}

	for i := range weekDays {
		fmt.Println(weekDays[i])
	}

	for i, day := range weekDays {
		fmt.Println(day, i)
	}

	// making a while loop

	val := 1

	for val < 10 {
		fmt.Println(val)
		val++
	}

	// continue
	val = 1

	for val < 10 {
		if val == 5 {
			val++
			continue // this is used to skip the below steps
		}
		fmt.Println(val, "11")
		val++
	}

}
