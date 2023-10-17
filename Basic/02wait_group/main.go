package main

import (
	"fmt"
	// "time"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup){
	// this will be executed when the function exits
	// this is used to notify the waitgroup that the goroutine is done
	// this will decrement the counter by 1
	defer wg.Done() 
	fmt.Println(s)
}


// the main function is intself is goroutine
// go sheduler is responsible for scheduling goroutines and managing the thread pool
func main(){

	// waitgroup is used to wait for the goroutines to finish
	// waitgroug variable is passed as a pointer to the goroutine
	var wg sync.WaitGroup

	words := [] string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"epsilon",
		"zeta",
		"eta",
		"theta",
		"iota",
		"kappa",
	}

	// Add the number of goroutines to the waitgroup
	wg.Add(len(words))

	// Print the words in a goroutine
	// This will be executed in a seperate thread
	// If we use time.Sleep to give time to the goroutine to execute
	// then the output will be in random order because the goroutines are spawned in order but excution is done by go sheduler
	for i, word := range words{
		go printSomething(fmt.Sprintf("%d: %s", i, word), &wg)
	}

	// wait for the goroutines to finish or counter to reach 0
	wg.Wait()

	// added so that it does not become negative, so for last function call we increment the counter
	wg.Add(1)
	//time.Sleep(1 * time.Second)
	
	printSomething("Second getting Printed", &wg)

}