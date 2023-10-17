package main

import (
	"fmt"
	"time"
)

func printSomething(s string){
	fmt.Println(s)
}


// the main function is intself is goroutine
// go sheduler is responsible for scheduling goroutines and managing the thread pool
func main(){

	// This is a goroutine 
	// This will be executed in a seperate thread
	// This does not get time to execute as the main thread is exited so fast
	go printSomething("First getting Printed")
	
	// we can use time.Sleep to give time to the goroutine to execute
	time.Sleep(1 * time.Second) 
	printSomething("Second getting Printed")

}