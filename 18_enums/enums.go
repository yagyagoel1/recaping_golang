package main

import "fmt"

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func changeOrderStatus(status OrderStatus) {

	fmt.Println("changing order statsu to", status)
}
func main() {
	changeOrderStatus(Received)
}
