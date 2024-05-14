package main

import (
	"fmt"
	"net/http"
	"sync"
)

func getStatusCode(wg *sync.WaitGroup, url string) {
	res, _ := http.Get(url)
	fmt.Printf("status code %v for url %v\n", res.StatusCode, url)
	wg.Done()

}
func main() {
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	urls := []string{
		"https://google.com",
		"https://go.dev",
		"https://github.com",
	}

	for _, url := range urls {
		wg.Add(1)
		go getStatusCode(wg, url)
	}

}
