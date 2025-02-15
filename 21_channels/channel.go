package main

import (
	"fmt"
)

//	func processNum(numChan chan int) {
//		for num := range numChan {
//			fmt.Println("processing number", num)
//			// time.Sleep(time.Second)
//		}
//	}
func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}
func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("processing...")
	done <- true
}
func main() {
	// messageChan := make(chan string)
	// messageChan <- "ping"//blocking
	// msg := <-messageChan
	// fmt.Println(msg)
	// var w sync.WaitGroup
	// defer w.Done()
	// w.Add(1)
	// numChan := make(chan int)
	// go processNum(numChan)
	// for {
	// 	numChan <- rand.Intn(100)
	// }
	// // w.Wait()
	// time.Sleep(time.Second)

	result := make(chan int)
	go sum(result, 4, 5)
	res := <-result //blocking that why no sleep
	fmt.Println(res)

	done := make(chan bool)

	go task(done)
	<-done

}
