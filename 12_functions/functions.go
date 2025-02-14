package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func getLanguages() (string, string, bool) {
	return "golang", "javascirpt", true
}
func processIt(fn func(a int) int) int {
	return fn(1)
}
func processF() func(a int) int {
	return func(a int) int {
		return a
	}
}
func main() {
	result := add(2, 4)
	fmt.Println("result", result)
	lang1, lang2, lang3 := getLanguages()
	fmt.Println(lang1)
	fmt.Println(lang2)
	fmt.Println(lang3)
	fn := func(a int) int {
		return 2
	}
	processIt(fn)
	fnn := processF()
}
