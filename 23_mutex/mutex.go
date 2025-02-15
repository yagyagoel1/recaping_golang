package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	p.mu.Lock()
	p.views += 1

	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()
}
func main() {
	myPost := post{views: 0}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)

	}
	wg.Wait()
	fmt.Println("", myPost.views)

}
