package main

import (
	"fmt"
)

func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func(done chan<- bool) { done <- true }(done)
	for email := range emailChan {
		// time.Sleep(time.Second)
		fmt.Println("send email to ", email)
	}
}
func main() {
	emailChan := make(chan string, 100)
	done := make(chan bool)
	// emailChan <- "1@example.com"
	// emailChan <- "2@ecample.com"

	// fmt.Println(<-emailChan)

	// fmt.Println(<-emailChan)
	go emailSender(emailChan, done)
	for i := 0; i < 100; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}
	close(emailChan)
	fmt.Println("done sending")
	<-done

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()
	go func() {
		chan2 <- "pong"
	}()
	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from cahn1", chan1Val)

		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)

		}
	}
}
