package main

import "fmt"

func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNumber", *num)
	var n int
	n = *num
	fmt.Println("n", n)
}
func main() {

	fmt.Println("")
	num := 1

	changeNum(&num)
	fmt.Println("after chanage num in main ", num)

}
