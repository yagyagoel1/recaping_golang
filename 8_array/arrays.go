package main

import "fmt"

func main() {
	var nums [4]int
	fmt.Println(len(nums))
	nums[0] = 5

	fmt.Println("get ", nums[0])
	fmt.Println("the whole array ", nums)
	var vals [4]bool
	fmt.Println(vals)
	var name [3]string
	name[0] = "golang"
	fmt.Println(name)

	num := []int{1, 2, 3}
	var nu []int
	nu = append(nu, 3)
	fmt.Println("num", num)
	fmt.Println("nu", nu)
	twodArray := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println("twodArray", twodArray)
}
