package main

import (
	"fmt"
	"time"
)

type customer struct {
	name    string
	phoneno string
}
type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}

func (o *Order) changeStatus(status string) {
	o.status = status
}
func (o *Order) getAmount() float32 {
	return o.amount
}
func newOrder(id string, amount float32, status string) *Order {
	order := Order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
	return &order
}
func main() {
	// order := newOrder("1", 50, "received")
	// fmt.Println("orderr", order)
	// fmt.Printf("%T", order)
	// order.changeStatus("confirmed")
	// fmt.Println(&order)
	// myOrder2 := Order{
	// 	id:        "3",
	// 	amount:    100,
	// 	status:    "delivered",
	// 	createdAt: time.Now(),
	// }
	// fmt.Println("order 2", myOrder2)

	language := struct {
		name   string
		isGood bool
	}{
		"golang", true}

	fmt.Println("", language)
	//struct embedding
	newOrder := Order{
		id:     "1",
		amount: 14,
		status: "received",
		customer: customer{
			name:    "yagya",
			phoneno: "phone"}}
	fmt.Println("", newOrder)
}
