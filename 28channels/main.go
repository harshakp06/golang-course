package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	myCh := make(chan int, 2)

	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)

	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChanelOpen := <-myCh

		fmt.Println(val)

		fmt.Println(isChanelOpen)
		// fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 5
		myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
