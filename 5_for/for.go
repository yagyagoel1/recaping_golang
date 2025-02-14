package main

import "fmt"

// for
func main() {
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//// infinite loop
	// for {
	// 	fmt.Println("loop")
	// }
	for i := range 3 {
		fmt.Println(i)
	}
}
