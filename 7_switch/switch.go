package main

import (
	"fmt"
	"time"
)

func main() {

	i := 5

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	default:
		fmt.Println("default")

	}

	// multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")

	}
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("its a integer ")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a booleann")
		default:
			fmt.Println("other", t)
		}
	}
	whoAmI(true)
}
