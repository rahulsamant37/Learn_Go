package main

import (
	"fmt"
	"time"
)

func main() {
	// Simple Switch
	// i := 3
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other")
	// }

	// Multiple Condition Switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
			fmt.Println("it's weekand")
	default:
			fmt.Println("it's workday")
	}
}