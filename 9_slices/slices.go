package main

import (
	"fmt"
	"slices"
)

func main() {
	// var nums []interface{}
	// fmt.Println(nums == nil)

	// nums = append(nums, 3)
	// nums = append(nums, "Fsa")
	// fmt.Println("num", nums)
	var nums = make([]int, 0, 5)
	fmt.Println("capcity", cap(nums))
	nums = append(nums, 1)
	nums = append(nums, 2)

	nums = append(nums, 3)
	fmt.Println("capcity", cap(nums))
	var nums2 = make([]int, len(nums))
	copy(nums2, nums)
	fmt.Println("nums2", nums2)
	fmt.Println("nums slcie", nums[0:2])
	fmt.Println(nums)
	fmt.Println("slices equal", slices.Equal(nums, nums2))
}
