package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // pointer

var signals = []string{"test"}

var mut sync.Mutex // pointer

func main() {
	// go greeter("Hello")
	// greeter("World")

	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	// fmt.Println(websitelist[2])
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()

	fmt.Println(signals)

}

// func greeter(s string) {

// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}

// }

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)

	defer wg.Done()

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {

		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
