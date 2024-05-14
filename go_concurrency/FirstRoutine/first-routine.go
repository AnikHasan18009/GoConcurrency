package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ragDao(junior string) {
	fmt.Println(junior + " ke rag dea ses!")
	time.Sleep(time.Second + time.Duration(rand.Intn(3)))
}
func main() {
	start := time.Now()
	juniors := []string{"anik", "mem", "momin"}
	for _, junior := range juniors {
		go ragDao(junior)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("agamikal abar hobe")
	fmt.Println(time.Since((start)))
}
