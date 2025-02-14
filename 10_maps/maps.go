package main

import (
	"fmt"
	"maps"
)

func main() {
	//create
	m := make(map[string]string)
	m["name"] = "golang"
	m["area"] = "backend"
	fmt.Println("", m["name"])
	fmt.Println(m["fd"] == "")
	delete(m, "Fd")
	// clear(m)
	ma := map[string]string{
		"value": "2"}

	fmt.Println(ma)
	fmt.Println("", m)
	v, ok := m["name"]
	fmt.Println(v)
	if ok {
		fmt.Println("all ok", ok)
	} else {
		fmt.Println("not ok ")
	}
	fmt.Printf("type: %T\n", ma)

	fmt.Println("equl", maps.Equal(m, ma))

}
