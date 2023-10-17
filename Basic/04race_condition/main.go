package main

import (
	"fmt"
	"sync"
)

var msg string

var wg sync.WaitGroup

func updateMeassage(s string) {
	msg = s
	defer wg.Done()
}

func main() {
	msg = "Hello"

	wg.Add(2)

	// here the race condition occurs
	// both are trying to update the same variable
	// so the last one to update will be the final value
	// To check the race is happening or not we can use the command
	// go run -race main.go
	go updateMeassage("First Goroutine")
	go updateMeassage("Second Goroutine")

	wg.Wait()

	fmt.Println(msg)
}
