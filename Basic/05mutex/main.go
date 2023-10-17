package main

import (
	"fmt"
	"sync"
)

var msg string

var wg sync.WaitGroup

func updateMeassage(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
	
}

func main() {
	msg = "Hello"

	// adding the mutex resolved the race condition issue.
	var mutex sync.Mutex

	wg.Add(2)
	go updateMeassage("First Goroutine", &mutex)
	go updateMeassage("Second Goroutine", &mutex)

	wg.Wait()

	fmt.Println(msg)
}
