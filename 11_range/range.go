package main

import "fmt"

func main() {
	nums := make([]int, 0, 5)
	nums = append(nums, 1)
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println("", nums[i])
	// }
	for key, nums := range nums {
		fmt.Println("", nums, key)
	}
	for i, c := range "golang" {
		fmt.Println("", i, c)
	}
}
