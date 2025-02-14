package main

import "fmt"

func printSlice[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)

	}

}

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

type stack[T comparable] struct {
	elements []T
}

func main() {
	// nums:=[]int{1,3,4}
	myStack := stack[int]{
		elements: []int{2, 4, 5},
	}
	fmt.Println("mystack", myStack)
	names := []string{"golang", "typescirpt"}
	printSlice(names)
}
