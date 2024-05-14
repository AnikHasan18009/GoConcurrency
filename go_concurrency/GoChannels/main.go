package main

import (
	"fmt"
	"time"
)

func ragDao(junior string, msg chan string) {
	time.Sleep(2 * time.Second)
	msg <- fmt.Sprint(junior + " ke rag dea ses!")

}
func main() {

	juniors := []string{"anik", "mem", "momin"}
	msg := make(chan string)
	for _, junior := range juniors {
		go ragDao(junior, msg)
	}
	i := 1
	for i < 4 {
		fmt.Println(<-msg)
		i++
	}

}
