package main

import "fmt"

func main() {
	name := "golang"
	name = "javascript"
	fmt.Println("Hello", name)
	const age = 30
	fmt.Println("My age is", age)
	// age = 40 // error
	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port, host)
}
